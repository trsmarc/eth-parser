package provider

import (
	"eth-parser/config"
	providerMocks "eth-parser/mocks/provider"
	storeMocks "eth-parser/mocks/txstore"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type EthIndexerTestSuite struct {
	suite.Suite
	config    config.Config
	ctrl      *gomock.Controller
	txStore   *storeMocks.MockTxStore
	ethClient *providerMocks.MockClient
	indexer   Indexer
}

func (suite *EthIndexerTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.config = config.Config{}
	suite.txStore = storeMocks.NewMockTxStore(suite.ctrl)
	suite.ethClient = providerMocks.NewMockClient(suite.ctrl)

	// TODO: Mock ethclient.Client and test NewEthIndexer
	// suite.indexer = NewEthIndexer(suite.config, suite.txStore, suite.ethClient)
}

func (suite *EthIndexerTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *EthIndexerTestSuite) TestGetCurrentBlock() {
	// TODO: Mock ethclient.Client and test GetCurrentBlock
}

func (suite *EthIndexerTestSuite) TestScanBlocks() {
	// TODO: Mock ethclient.Client and test ScanBlocks
}

func TestEthIndexerTestSuite(t *testing.T) {
	suite.Run(t, new(EthIndexerTestSuite))
}
