package pkg

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func getWalletClientPair(coin CoinSymbol) (IWallet, IClient, error) {

	switch coin {
	case ETH:
		wallet, err := GetWalletManager().GetWallet(coin)
		if err != nil {
			return nil, nil, err
		}
		client, err := GetClientStorage().GetClient(coin, wallet.GetNetwork())
		if err != nil {
			return nil, nil, err
		}
		return wallet, client, nil
	default:
		return nil, nil, UnknownCoinSymbolError
	}
}

func GetBalance(coin CoinSymbol) (*big.Float, error) {
	w, c, err := getWalletClientPair(coin)
	if err != nil {
		return nil, err
	}

	tb := big.NewFloat(0)
	for _, pk := range w.getPrivateKeys() {
		b, err := c.GetBalance(crypto.PubkeyToAddress(pk.PublicKey).Bytes())
		if err != nil {
			return nil, err
		}
		tb = tb.Add(tb, b)
	}
	return tb, nil
}

func MakeTransaction(coin CoinSymbol, to string, amount float64) error {
	w, c, err := getWalletClientPair(coin)
	if err != nil {
		return err
	}

	pks := w.getPrivateKeys()
	addr := common.HexToAddress(to)
	c.SendTransaction(pks, &addr, amount)
	return nil
}
