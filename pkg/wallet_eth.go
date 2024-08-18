package pkg

import (
	"crypto/ecdsa"
	"fmt"
	"sync/atomic"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

type ethw struct {
	mnemonic    string
	supportWord string
	network     string
	privateKeys []*ecdsa.PrivateKey
	accountId   int
	//
	masterKey    *hdkeychain.ExtendedKey
	sym          CoinSymbol
	name         CoinName
	addressIndex atomic.Int64
}

func GetEthWallet(config walletConfig) (*ethw, error) {
	seed := bip39.NewSeed(config.mnemonic, config.supportWord)
	k, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	pks := make([]*ecdsa.PrivateKey, len(config.privateKeys))
	for i, k := range config.privateKeys {
		pk, err := crypto.HexToECDSA(k)
		if err != nil {
			return nil, err
		}
		pks[i] = pk
	}
	return &ethw{
		mnemonic:     config.mnemonic,
		supportWord:  config.supportWord,
		network:      config.network,
		privateKeys:  pks,
		masterKey:    k,
		addressIndex: atomic.Int64{},
		accountId:    0,
		sym:          ETH,
		name:         ETH_N,
	}, nil
}

func (e *ethw) GetSymbol() CoinSymbol {
	return e.sym
}

func (e *ethw) GetName() CoinName {
	return e.name
}

func (e *ethw) NewPrivateKey() (string, error) {
	// TODO: if we need internal network keys, implement new one,
	// and set to 1
	change := 0
	childKey, err := deriveKey(e.masterKey, int(ETH_T), e.accountId, change, int(e.addressIndex.Load()))
	if err != nil {
		return "", err
	}
	btcecPK, err := childKey.ECPrivKey()
	if err != nil {
		return "", err
	}
	key := fmt.Sprintf("%x", btcecPK.Serialize())
	e.addressIndex.Add(1)
	return key, nil
}

// figure out a way to separate network, and the relation between wallet and client
func (e *ethw) GetNetwork() string {
	return e.network
}

func (e *ethw) getPrivateKeys() []*ecdsa.PrivateKey {
	return e.privateKeys
}
