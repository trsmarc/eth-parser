package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Transaction struct {
	// transaction hash
	Hash string
	// block number
	Block int
	// sender address
	From string
	// receiver address
	To string
	// amount
	Value *big.Int
}

type EthClient interface {
	*ethclient.Client
}
