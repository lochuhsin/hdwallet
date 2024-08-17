package pkg

import "fmt"

type IWallet interface {
	GetName() CoinName
	GetSymbol() CoinSymbol
}

type walletManager struct {
	wallets map[CoinSymbol]IWallet
}

func (w *walletManager) GetWallet(type_ CoinSymbol, mnConfig MnemonicConfig) (IWallet, error) {
	cli, ok := w.wallets[type_]
	if !ok {
		client, err := w.initWallet(type_, mnConfig)
		w.wallets[type_] = client
		if err != nil {
			return nil, err
		}
		cli = client
	}
	return cli, nil
}

func (w walletManager) initWallet(type_ CoinSymbol, mnConfig MnemonicConfig) (IWallet, error) {
	switch type_ {
	case ETH:
		return getEthWallet(mnConfig)
	default:
		return nil, fmt.Errorf("Unsupported CoinType: %v", type_)
	}
}

func newWalletStorage() *walletManager {
	return &walletManager{
		wallets: make(map[CoinSymbol]IWallet),
	}
}
