package pkg

import "github.com/tyler-smith/go-bip39"

func NewMnemonic() (string, error) {
	entropy, _ := bip39.NewEntropy(256)
	return bip39.NewMnemonic(entropy)
}
