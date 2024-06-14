package config

type Config struct {
	Port     string `envconfig:"PORT" default:"8080"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	Indexer  IndexerConfig
}

type IndexerConfig struct {
	Endpoint          string `envconfig:"ENDPOINT" default:"https://mainnet.infura.io/v3/12bd8c6b2e5e400b8cf17b6664320653"`
	StartBlock        uint64 `envconfig:"START_BLOCK" default:"20083813"`
	ConcurrencyLimit  int    `envconfig:"CONCURRENCY_LIMIT" default:"10"`
	BatchSize         uint64 `envconfig:"BATCH_SIZE" default:"50"`
	RateLimitInterval int    `envconfig:"RATE_LIMIT_INTERVAL" default:"1"`
	NewBlockInterval  int    `envconfig:"NEW_BLOCK_INTERVAL" default:"15"`
}
