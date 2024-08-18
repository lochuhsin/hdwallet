package pkg

import "github.com/tyler-smith/go-bip39"

func NewMnemonic() (MnemonicConfig, error) {
	entropy, _ := bip39.NewEntropy(256)
	mn, err := bip39.NewMnemonic(entropy)
	return MnemonicConfig{MN: mn}, err
}
