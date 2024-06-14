package store

import (
	"errors"
	"eth-parser/common"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

type MemoryStore struct {
	store map[string][]common.Transaction
}

func NewMemoryStore() TxStore {
	store := make(map[string][]common.Transaction)
	return &MemoryStore{
		store,
	}
}

func addressCheck(address string) bool {
	addr := ethCommon.HexToAddress(address)
	return ethCommon.IsHexAddress(address) && addr.Hex() == address
}

func (m *MemoryStore) Add(address string) error {
	if !addressCheck(address) {
		return errors.New("invalid address")
	}

	_, ok := m.store[address]
	if ok {
		return errors.New("address already subscribed")
	}

	m.store[address] = make([]common.Transaction, 0)

	return nil
}

func (m *MemoryStore) Append(address string, tx common.Transaction) {
	m.store[address] = append(m.store[address], tx)
}

func (m *MemoryStore) List(address string) ([]common.Transaction, error) {
	if !addressCheck(address) {
		return []common.Transaction{}, errors.New("invalid address")
	}

	return m.store[address], nil
}

func (m *MemoryStore) Keys() []string {
	keys := make([]string, 0, len(m.store))
	for k := range m.store {
		keys = append(keys, k)
	}

	return keys
}
