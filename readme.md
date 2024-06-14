# eth-parser

## Table of Contents
- [eth-parser](#eth-parser)
  - [Table of Contents](#table-of-contents)
  - [About the project](#about-the-project)
  - [API documentation](#api-documentation)
  - [Live DEMO](#live-demo)
  - [Local development](#local-development)
    - [Running from source](#running-from-source)
      - [Prerequisites](#prerequisites)
    - [Project Layout](#project-layout)

## About the project

The purpose of the project is to provides a RESTful API to interact with the transaction on the network by iterate through Ethereum blocks and transactions, store parsed transaction for subscribed addresses into the in-memory store.

## API documentation
For full API document please checkout [this section](https://github.com/trsmarc/eth-parser/tree/main/apidoc)

## Live DEMO
API server was deployed to GCP Cloud Run, you can access it via this [link](#live-demo)

- Subscribe to an address 
- List collected transactions of an address
- Get current fetched block number
- List all addresses that have been subscribed

*Note* The live demo is using concurrency limit of 10, batch size of 50 to prevent the service from being rate limited by the blockchain provider. Please try running locally with different configuration for better performance.

## Local development

### Running from source

#### Prerequisites
- Go 1.22.3

Override environment variables is optional, you can use the default values
```bash
# default values
export PORT=8080
export LOG_LEVEL=info
export ENDPOINT="https://mainnet.infura.io/v3/12bd8c6b2e5e400b8cf17b6664320653"
export START_BLOCK=20083813
export CONCURRENCY_LIMIT=10
export BATCH_SIZE=50
export RATE_LIMIT_INTERVAL=1
export NEW_BLOCK_INTERVAL=15
```

Run the server:
```bash
go run main.go
```

### Project Layout

```
.
├── common
│   └── ...
├── config
│   └── ...
├── provider
│   └── ...
├── server
│   └── ...
└── txstore
    └── ...
```

`common` - Shared utilities, data structures, and helper functions.

`config` - Application settings, connections, keys.

`provider` - Interfaces and implementation for different blockchain provider.

`server` - HTTP server setup, handler, and API routes.

`txstore` - Transaction data storage and retrieval logic.

