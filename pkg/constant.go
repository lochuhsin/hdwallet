package pkg

const FILE_PERMISSIONS = 0666

type CoinSymbol string
type CoinType int
type CoinName string

/**
 * The registered coin type follows the slip-0044
 * https://github.com/satoshilabs/slips/blob/master/slip-0044.md
 */

const (
	ETH CoinSymbol = "ETH"
	ETC CoinSymbol = "ETC"
)

const (
	ETH_T CoinType = 60
	ETH_C CoinType = 61
)

const (
	ETH_N CoinName = "Ether"
	ETC_N CoinName = "Ether Classic"
)

func CoinSelector(symbol string) (CoinSymbol, error) {
	c := CoinSymbol(symbol)

	switch c {
	case ETH:
		return ETH, nil
	case ETC:
		return ETC, nil
	default:
		return "", UnknownCoinSymbolError
	}
}
