package reports

type CashAppTransaction struct {
	TransactionID        string
	Date                 string
	TransactionType      string
	Currency             string
	Amount               string
	Fee                  string
	NetAmount            string
	AssetType            string
	AssetPrice           string
	AssetAmount          string
	Status               string
	Notes                string
	NameOfSenderReceiver string
	Account              string
}

func (cashAppTransaction *CashAppTransaction) GetCounterparty() string {
	return cashAppTransaction.Notes
}

func NewCashAppTransaction(record []string) CashAppTransaction {
	return CashAppTransaction{
		TransactionID:        record[0],
		Date:                 record[1],
		TransactionType:      record[2],
		Currency:             record[3],
		Amount:               record[4],
		Fee:                  record[5],
		NetAmount:            record[6],
		AssetType:            record[7],
		AssetPrice:           record[8],
		AssetAmount:          record[9],
		Status:               record[10],
		Notes:                record[11],
		NameOfSenderReceiver: record[12],
		Account:              record[13],
	}
}
