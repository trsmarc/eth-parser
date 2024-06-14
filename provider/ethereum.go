package provider

import (
	"context"
	"eth-parser/common"
	"eth-parser/config"
	store "eth-parser/txstore"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthIndexer struct {
	store             store.TxStore
	config            config.Config
	client            *ethclient.Client
	nextBlock         uint64
	chainID           *big.Int
	concurrencyLimit  int           // number of concurrent requests
	batchSize         uint64        // number of blocks to process in a batch
	rateLimitInterval time.Duration // rate limit interval
	newBlockInterval  time.Duration // delay to wait for new block
}

func NewEthIndexer(cfg config.Config, txStore store.TxStore, client *ethclient.Client) Indexer {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return &EthIndexer{
		config:            cfg,
		client:            client,
		chainID:           chainID,
		store:             txStore,
		nextBlock:         cfg.Indexer.StartBlock,
		concurrencyLimit:  cfg.Indexer.ConcurrencyLimit,
		batchSize:         cfg.Indexer.BatchSize,
		rateLimitInterval: time.Duration(cfg.Indexer.RateLimitInterval) * time.Second,
		newBlockInterval:  time.Duration(cfg.Indexer.NewBlockInterval) * time.Second,
	}
}

// GetCurrentBlock - get the last scanned block number
func (e *EthIndexer) GetCurrentBlock() uint64 {
	return e.nextBlock
}

// ScanBlock - scan the Ethereum blockchain for a new block and process it with concurrency
func (e *EthIndexer) ScanBlocks() {
	sem := make(chan struct{}, e.concurrencyLimit)
	startBlock := e.nextBlock

	for blockNum := startBlock; ; blockNum += e.batchSize {
		latestBlock, err := e.getLatestBlock()
		if err != nil {
			common.Logger.Error().Err(err).Msg("Error getting latest block")
			continue
		}

		if e.nextBlock > latestBlock {
			common.Logger.Warn().Msg("Latest block behind current block, Waiting for new block ...")
			time.Sleep(e.newBlockInterval)
			continue
		}

		sem <- struct{}{} // Acquire a semaphore slot

		go func(blockNum uint64) {
			defer func() { <-sem }() // Release the semaphore slot

			for i := blockNum; i < blockNum+e.batchSize; i++ {
				e.nextBlock = i
				logger := common.Logger.With().Uint64("block", i).Logger()
				logger.Info().Msg("Fetching block")

				block, err := e.client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
				if err != nil {
					logger.Error().Err(err).Msg("Error fetching block")
					continue
				}

				// process the block
				err = e.processBlock(block)
				if err != nil {
					logger.Error().Err(err).Msg("Error processing block")
				}
			}
		}(blockNum)

		time.Sleep(e.rateLimitInterval)
	}
}

// processBlock - process the block and store the data
func (e *EthIndexer) processBlock(block *types.Block) error {
	logger := common.Logger.With().Uint64("block", block.Number().Uint64()).Logger()
	logger.Debug().Msgf("Processing block")
	addrs := e.getSubscribedAddress()

	txs := block.Transactions()
	if txs == nil {
		return nil
	}

	txMatchCount := 0

	// process the block
	for _, tx := range txs {
		var toAddr, fromAddr string
		if tx == nil {
			continue
		}

		if tx.To() != nil {
			toAddr = tx.To().Hex()
		}

		sender, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			logger.Warn().Str("txHash", tx.Hash().String()).Err(err).Msg("Error getting sender address")
		} else {
			fromAddr = sender.Hex()
		}

		txToAdd := common.Transaction{
			Hash:  tx.Hash().Hex(),
			Block: int(block.Number().Int64()),
			From:  fromAddr,
			To:    toAddr,
			Value: tx.Value(),
		}

		// check if the address is subscribed, if so add the transaction to the store
		if addrs[toAddr] {
			txMatchCount++
			e.store.Append(toAddr, txToAdd)
		}

		if addrs[fromAddr] {
			txMatchCount++
			e.store.Append(fromAddr, txToAdd)
		}
	}

	logger.Debug().Msgf("Total tx in block: %d, Total tx matched: %d", len(txs), txMatchCount)

	return nil
}

// getLatestBlock - get the latest block number from the Ethereum blockchain
func (e *EthIndexer) getLatestBlock() (uint64, error) {
	header, err := e.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	return header.Number.Uint64(), nil
}

// getSubscribedAddress - get all the subscribed addresses
func (e *EthIndexer) getSubscribedAddress() map[string]bool {
	addrs := make(map[string]bool)
	keys := e.store.Keys()
	for _, key := range keys {
		addrs[key] = true
	}

	return addrs
}
