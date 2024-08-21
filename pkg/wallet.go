package pkg

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"sync"
)

var walletM *walletManager
var wOnce sync.Once // guard the initialization of storage

type IWallet interface {
	GetName() CoinName
	GetSymbol() CoinSymbol
	GetNetwork() string
	NewPrivateKey() (string, error)
	getPrivateKeys() []*ecdsa.PrivateKey
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
	wallet, err := walletSelector(type_, opts...)
	if err != nil {
		return nil, err
	}
	w.wallets[type_] = wallet
	return wallet, nil
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
	password    string
	accountId   int
	privateKeys []string
	network     string
}

func newWalletConfig() walletConfig {
	return walletConfig{
		privateKeys: []string{},
		mnemonic:    "",
		password:    "",
		accountId:   0,
		network:     "",
	}
}

type WalletOpt func(conf *walletConfig)

func SetMnemonic(mnemonic string) WalletOpt {
	return func(w *walletConfig) {
		w.mnemonic = mnemonic
	}
}

func SetPassword(password string) WalletOpt {
	return func(w *walletConfig) {
		w.password = password
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

func SetNetwork(network string) WalletOpt {
	return func(w *walletConfig) {
		w.network = network
	}
}

func walletSelector(type_ CoinSymbol, opts ...WalletOpt) (IWallet, error) {
	wConfig := newWalletConfig()
	for _, opt := range opts {
		opt(&wConfig)
	}

	switch type_ {
	case ETH:
		return GetEthWallet(wConfig)
	default:
		return nil, fmt.Errorf("Unsupported CoinType: %v", type_)
	}
}
