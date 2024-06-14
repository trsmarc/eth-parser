package store

import (
	"eth-parser/common"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MemoryTxStoreTestSuite struct {
	suite.Suite
	store TxStore
}

func (suite *MemoryTxStoreTestSuite) SetupTest() {
	suite.store = NewMemoryStore()
}

func (suite *MemoryTxStoreTestSuite) TestAdd() {
	tests := []struct {
		title   string
		address string
		err     bool
	}{
		{
			title:   "valid address",
			address: "0x1234567890123456789012345678901234567890",
			err:     false,
		},
		{
			title:   "invalid address",
			address: "invalid",
			err:     true,
		},
	}

	for _, test := range tests {
		err := suite.store.Add(test.address)
		if test.err {
			suite.Error(err)
		} else {
			suite.NoError(err)
		}
	}
}

func (suite *MemoryTxStoreTestSuite) TestAddDuplicate() {
	suite.NoError(suite.store.Add("0x1234567890123456789012345678901234567890"))
	err := suite.store.Add("0x1234567890123456789012345678901234567890")
	suite.Error(err)
}

func (suite *MemoryTxStoreTestSuite) TestAppend() {
	tests := []struct {
		title   string
		address string
		tx      common.Transaction
	}{
		{
			title:   "valid transaction",
			address: "0x1234567890123456789012345678901234567890",
			tx: common.Transaction{
				Hash: "0x1234567890123456789012345678901234567890123456789012345678901234",
			},
		},
	}

	for _, test := range tests {
		suite.NoError(suite.store.Add(test.address))
		suite.store.Append(test.address, test.tx)

		txs, err := suite.store.List(test.address)
		suite.NoError(err)
		suite.Equal([]common.Transaction{test.tx}, txs)
	}
}

func (suite *MemoryTxStoreTestSuite) TestList() {
	tests := []struct {
		title   string
		address string
		err     bool
	}{
		{
			title:   "valid address",
			address: "0x1234567890123456789012345678901234567890",
			err:     false,
		},
		{
			title:   "invalid address",
			address: "invalid",
			err:     true,
		},
	}

	for _, test := range tests {
		_, err := suite.store.List(test.address)
		if test.err {
			suite.Error(err)
		} else {
			suite.NoError(err)
		}
	}
}

func (suite *MemoryTxStoreTestSuite) TestKeys() {
	suite.store.Add("0x1234567890123456789012345678901234567890")
	suite.store.Add("0x1234567890123456789012345678901234567891")

	keys := suite.store.Keys()
	suite.Len(keys, 2)
}

func TestMemoryTxStoreTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryTxStoreTestSuite))
}
