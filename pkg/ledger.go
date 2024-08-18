package pkg

import (
	"math/big"

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

func GetBalance(coin CoinSymbol) (*big.Int, error) {
	wallet, client, err := getWalletClientPair(coin)
	if err != nil {
		return nil, err
	}

	balance := big.NewInt(0)
	for _, pk := range wallet.getPrivateKeys() {
		address := crypto.PubkeyToAddress(pk.PublicKey)

		b, err := client.GetBalance(address.Bytes())
		if err != nil {
			return nil, err
		}
		balance = balance.Add(balance, b)
	}
	return balance, nil
}
