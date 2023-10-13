package reports

type Report interface {
	ParseFileRecords(filePath string) error
	GetTransactions() []Transaction
}

type Transaction interface {
	GetDate() string
	GetAmount() float32
	GetDescription() string
	GetCounterparty() string
}
