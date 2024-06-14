package provider

import "github.com/ethereum/go-ethereum"

type Indexer interface {
	ScanBlocks()
	GetCurrentBlock() uint64
}

// Client exposes the methods provided by the Ethereum RPC client.
type EthClient interface {
	ethereum.BlockNumberReader
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.GasPricer1559
	ethereum.FeeHistoryReader
	ethereum.LogFilterer
	ethereum.PendingStateReader
	ethereum.PendingContractCaller
	ethereum.TransactionReader
	ethereum.TransactionSender
	ethereum.ChainIDReader
}
