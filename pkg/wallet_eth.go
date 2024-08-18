package pkg

import (
	"crypto/ecdsa"
	"sync/atomic"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type ethw struct {
	mnemonic     string
	supportWord  string
	masterKey    *hdkeychain.ExtendedKey
	sym          CoinSymbol
	name         CoinName
	privateKeys  []*ecdsa.PrivateKey
	accountId    int
	addressIndex atomic.Int64
}

func getEthWallet(config walletConfig) (*ethw, error) {
	seed := bip39.NewSeed(config.mnemonic, config.supportWord)
	k, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	return &ethw{
		mnemonic:     config.mnemonic,
		supportWord:  config.supportWord,
		privateKeys:  []*ecdsa.PrivateKey{},
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
	key := EncodePrivateKeyToHex(btcecPK.ToECDSA())
	e.addressIndex.Add(1)
	return key, nil
}
