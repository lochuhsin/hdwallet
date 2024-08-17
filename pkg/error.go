package pkg

import "errors"

var UnknownCoinSymbolError = errors.New("unknown coin symbol")

var InvalidHostError = errors.New("Unable to connect to host")
