package main

import (
	"eth-parser/config"
	"eth-parser/provider"
	"eth-parser/server"
	store "eth-parser/txstore"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	// read the config from environment
	var config config.Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}

	// initialize the store
	memStore := store.NewMemoryStore()

	client, err := ethclient.Dial(config.Indexer.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// initialize the Ethereum indexer
	ethIndexer := provider.NewEthIndexer(config, memStore, client)
	go ethIndexer.ScanBlocks()

	// start the server
	s := server.NewServer(ethIndexer, memStore)
	s.StartServer()
}
