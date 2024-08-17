package pkg

import (
	"crypto/ecdsa"
	"errors"
	"sync/atomic"

	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type ethWallet struct {
	mnemonic     MnemonicConfig
	masterKey    *hdkeychain.ExtendedKey
	internalKeys []*ecdsa.PrivateKey
	externalKeys []*ecdsa.PrivateKey
	sym          CoinSymbol
	name         CoinName
	addressIndex atomic.Int64
	accountId    int
}

func getEthWallet(config MnemonicConfig) (*ethWallet, error) {
	seed := bip39.NewSeed(config.MN, config.SupportWord)
	k, err := hdkeychain.NewMaster(seed, nil)
	if err != nil {
		return nil, err
	}
	return &ethWallet{
		mnemonic:     config,
		masterKey:    k,
		internalKeys: []*ecdsa.PrivateKey{},
		externalKeys: []*ecdsa.PrivateKey{},
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

func (e *ethWallet) GenExternalKey() error {
	k, err := e.genNewKey(0)
	e.externalKeys = append(e.externalKeys, k)
	if err != nil {
		return err
	}
	e.addressIndex.Add(1)
	return nil
}

func (e *ethWallet) GenInternalKey() error {
	k, err := e.genNewKey(1)
	e.internalKeys = append(e.internalKeys, k)
	if err != nil {
		return err
	}
	e.addressIndex.Add(1)
	return nil
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
