package pkg

import (
	"crypto/ecdsa"
	"errors"
	"sync/atomic"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type ethWallet struct {
	mnemonic     MnemonicConfig
	masterKey    *hdkeychain.ExtendedKey
	sym          CoinSymbol
	name         CoinName
	addressIndex atomic.Int64
	privateKeys  []*ecdsa.PrivateKey
	accountId    int
}

func getEthWallet(config MnemonicConfig) (*ethWallet, error) {
	seed := bip39.NewSeed(config.MN, config.SupportWord)
	k, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	return &ethWallet{
		mnemonic:     config,
		masterKey:    k,
		sym:          ETH,
		name:         ETH_N,
		addressIndex: atomic.Int64{},
		accountId:    0,
	}, nil
}

func (e *ethWallet) GetSymbol() CoinSymbol {
	return e.sym
}

func (e *ethWallet) GetName() CoinName {
	return e.name
}

func (e *ethWallet) GenExternalKey() (*ecdsa.PrivateKey, error) {
	k, err := e.genNewKey(0)
	if err != nil {
		return nil, err
	}
	e.addressIndex.Add(1)
	return k, nil
}

func (e *ethWallet) GenInternalKey() (*ecdsa.PrivateKey, error) {
	k, err := e.genNewKey(1)
	if err != nil {
		return nil, err
	}
	e.addressIndex.Add(1)
	return k, nil
}

func (e *ethWallet) genNewKey(change int) (*ecdsa.PrivateKey, error) {
	if change != 0 && change != 1 {
		return nil, errors.New("invalid change number, must be 0 or 1")
	}
	childKey, err := deriveKey(e.masterKey, int(ETH_T), e.accountId, 1, int(e.addressIndex.Load()))
	if err != nil {
		return nil, err
	}
	btcecPK, err := childKey.ECPrivKey()
	if err != nil {
		return nil, err
	}
	return btcecPK.ToECDSA(), nil
}
