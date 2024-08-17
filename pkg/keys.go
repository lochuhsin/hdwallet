package pkg

import "github.com/btcsuite/btcutil/hdkeychain"

const PURPOSE = 44 // bip44

func deriveKey(master *hdkeychain.ExtendedKey, coinType int, accountId int, change int, addrIndex int) (*hdkeychain.ExtendedKey, error) {
	/**
	 * 44 / coin type / account id / change / address index /
	 */

	path := []int{PURPOSE, coinType, accountId, change, addrIndex}

	key := master
	for _, p := range path {
		k, err := key.Child(hdkeychain.HardenedKeyStart + uint32(p))
		if err != nil {
			return nil, err
		}
		key = k
	}
	return key, nil
}
