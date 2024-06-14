package config

type Config struct {
	Port     string `default:"8080"`
	LogLevel string `default:"info"`
	Indexer  IndexerConfig
}

type IndexerConfig struct {
	Endpoint          string `default:"https://mainnet.infura.io/v3/12bd8c6b2e5e400b8cf17b6664320653"`
	StartBlock        uint64 `default:"20083813"`
	ConcurrencyLimit  int    `default:"10"`
	BatchSize         uint64 `default:"50"`
	RateLimitInterval int    `default:"1"`
	NewBlockInterval  int    `default:"15"`
}
