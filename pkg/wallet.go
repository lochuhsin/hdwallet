package pkg

import (
	"errors"
	"fmt"
	"sync"
)

var walletM *walletManager
var wOnce sync.Once // guard the initialization of storage

type IWallet interface {
	GetName() CoinName
	GetSymbol() CoinSymbol
	NewPrivateKey() (string, error)
}

type walletManager struct {
	wallets map[CoinSymbol]IWallet
}

func (w *walletManager) GetWallet(type_ CoinSymbol) (IWallet, error) {
	cw, ok := w.wallets[type_]
	if !ok {
		return nil, errors.New("No wallet found with current coin type")
	}
	return cw, nil
}

func (w *walletManager) NewWallet(type_ CoinSymbol, opts ...WalletOpt) (IWallet, error) {
	client, err := walletSelector(type_, opts...)
	if err != nil {
		return nil, err
	}
	w.wallets[type_] = client
	return client, nil
}

func newWalletStorage() *walletManager {
	return &walletManager{
		wallets: make(map[CoinSymbol]IWallet),
	}
}

func InitWalletManager() {
	wOnce.Do(func() {
		walletM = newWalletStorage()
	})
}

func GetWalletManager() *walletManager {
	return walletM
}

type walletConfig struct {
	mnemonic    string
	supportWord string
	accountId   int
	privateKeys []string
}

func newWalletConfig() walletConfig {
	return walletConfig{
		privateKeys: []string{},
		mnemonic:    "",
		supportWord: "",
		accountId:   0,
	}
}

type WalletOpt func(conf *walletConfig)

func SetMnemonic(mnemonic string) WalletOpt {
	return func(w *walletConfig) {
		w.mnemonic = mnemonic
	}
}

func SetSupportWord(supportWord string) WalletOpt {
	return func(w *walletConfig) {
		w.supportWord = supportWord
	}
}

func SetPrivateKeys(keys []string) WalletOpt {
	return func(w *walletConfig) {
		w.privateKeys = keys
	}
}

func SetAccountId(id int) WalletOpt {
	return func(w *walletConfig) {
		w.accountId = id
	}
}

func walletSelector(type_ CoinSymbol, opts ...WalletOpt) (IWallet, error) {
	wConfig := newWalletConfig()
	for _, opt := range opts {
		opt(&wConfig)
	}

	switch type_ {
	case ETH:
		return getEthWallet(wConfig)
	default:
		return nil, fmt.Errorf("Unsupported CoinType: %v", type_)
	}
}
