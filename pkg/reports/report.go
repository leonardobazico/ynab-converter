package reports

import (
	"time"
)

type Report interface {
	ParseFileRecords(filePath string) error
	GetTransactions() []Transaction
}

type Transaction interface {
	GetCounterparty() string
	GetDescription() string
	GetAmount() (float32, error)
	GetDatetime() (*time.Time, error)
}
