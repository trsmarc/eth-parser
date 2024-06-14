package server

import (
	"bytes"
	"encoding/json"
	"eth-parser/common"
	indexerMocks "eth-parser/mocks/provider"
	storeMocks "eth-parser/mocks/txstore"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	indexer *indexerMocks.MockIndexer
	txStore *storeMocks.MockTxStore
	server  Server
}

func (suite *ServerTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.indexer = indexerMocks.NewMockIndexer(suite.ctrl)
	suite.txStore = storeMocks.NewMockTxStore(suite.ctrl)
	suite.server = NewServer(suite.indexer, suite.txStore)
}

func (suite *ServerTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *ServerTestSuite) TestSubscribe() {
	tests := []struct {
		description   string
		requestBody   SubscribeRequest
		mockFunc      func()
		expectedBody  string
		expectedError bool
	}{
		{
			description: "Successful subscription",
			requestBody: SubscribeRequest{
				Address: "0x123",
			},
			mockFunc: func() {
				suite.txStore.EXPECT().Add("0x123").Return(nil)
			},
			expectedBody:  `{"message":"Subscribed successfully"}`,
			expectedError: false,
		},
	}

	for _, test := range tests {
		ts := httptest.NewServer(http.HandlerFunc(suite.server.subscribe))
		defer ts.Close()

		test.mockFunc()

		path := "/subscribe"
		jsonData, err := json.Marshal(&test.requestBody)
		suite.NoError(err)

		req, _ := http.NewRequest(http.MethodPost, ts.URL+path, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		suite.Nil(err)
		suite.NotNil(resp)

		body, err := io.ReadAll(resp.Body)
		suite.Nil(err)
		suite.Contains(string(body), test.expectedBody)
		suite.Equal(http.StatusOK, resp.StatusCode)
		resp.Body.Close()
	}
}

func (suite *ServerTestSuite) TestGetCurrentBlock() {
	tests := []struct {
		description   string
		mockFunc      func()
		expectedBody  string
		expectedError bool
	}{
		{
			description: "Successful get current block",
			mockFunc: func() {
				suite.indexer.EXPECT().GetCurrentBlock().Return(uint64(123))
			},
			expectedBody:  `{"blockNumber":123}`,
			expectedError: false,
		},
	}

	for _, test := range tests {
		ts := httptest.NewServer(http.HandlerFunc(suite.server.getCurrentBlock))
		defer ts.Close()

		suite.Run(test.description, func() {
			test.mockFunc()
		})

		path := "/block"
		resp, err := http.Get(ts.URL + path)
		if err != nil {
			log.Fatal(err)
		}

		suite.Nil(err)
		suite.NotNil(resp)

		body, err := io.ReadAll(resp.Body)
		suite.Nil(err)
		suite.Contains(string(body), test.expectedBody)
		suite.Equal(http.StatusOK, resp.StatusCode)
		resp.Body.Close()
	}
}

func (suite *ServerTestSuite) TestGetSubscribedAddress() {
	tests := []struct {
		description   string
		mockFunc      func()
		expectedBody  string
		expectedError bool
	}{
		{
			description: "Successful get subscribed address",
			mockFunc: func() {
				suite.txStore.EXPECT().Keys().Return([]string{"0x123"})
			},
			expectedBody:  `{"addresses":["0x123"],"total":1}`,
			expectedError: false,
		},
	}

	for _, test := range tests {
		ts := httptest.NewServer(http.HandlerFunc(suite.server.getSubscribedAddress))
		defer ts.Close()

		suite.Run(test.description, func() {
			test.mockFunc()
		})

		path := "/subscriber"
		resp, err := http.Get(ts.URL + path)
		if err != nil {
			log.Fatal(err)
		}

		suite.Nil(err)
		suite.NotNil(resp)

		body, err := io.ReadAll(resp.Body)
		suite.Nil(err)
		suite.Contains(string(body), test.expectedBody)
		suite.Equal(http.StatusOK, resp.StatusCode)
		resp.Body.Close()
	}
}

func (suite *ServerTestSuite) TestGetTransactions() {
	tests := []struct {
		description   string
		requestParam  string
		mockFunc      func()
		expectedBody  string
		expectedError bool
	}{
		{
			description:  "Successful get transactions",
			requestParam: "address=0x123",
			mockFunc: func() {
				suite.txStore.EXPECT().List("0x123").Return([]common.Transaction{
					{
						Hash:  "0x123",
						From:  "0x456",
						To:    "0x789",
						Value: big.NewInt(1e18),
					},
				}, nil)
			},
			expectedBody:  `{"total":1,"transactions":[{"Hash":"0x123","Block":0,"From":"0x456","To":"0x789","Value":1000000000000000000}]}`,
			expectedError: false,
		},
	}

	for _, test := range tests {
		ts := httptest.NewServer(http.HandlerFunc(suite.server.getTransactions))
		defer ts.Close()

		test.mockFunc()

		path := "/transactions?" + test.requestParam
		req, _ := http.NewRequest(http.MethodGet, ts.URL+path, nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		suite.Nil(err)
		suite.NotNil(resp)

		body, err := io.ReadAll(resp.Body)
		suite.Nil(err)
		suite.Contains(string(body), test.expectedBody)
		suite.Equal(http.StatusOK, resp.StatusCode)
		resp.Body.Close()
	}
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
