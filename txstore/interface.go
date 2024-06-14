package store

import "eth-parser/common"

type TxStore interface {
	Add(address string) error
	Append(address string, tx common.Transaction)
	List(address string) ([]common.Transaction, error)
	Keys() []string
}
