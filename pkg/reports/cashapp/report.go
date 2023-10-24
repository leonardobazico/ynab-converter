package cashapp

import (
	"fmt"

	"cash2ynab/internal/file"
)

type Report struct {
	fileRecordsGetter file.RecordsGetter
	transactions      []Transaction
}

func (cashAppReport *Report) ParseFileRecords(filePath string) error {
	records, err := cashAppReport.fileRecordsGetter.GetRecordsFrom(filePath)
	if err != nil {
		return fmt.Errorf("fail to get records from file: %w", err)
	}

	for _, record := range records {
		cashAppReport.transactions = append(
			cashAppReport.transactions,
			NewCashAppTransaction(record),
		)
	}

	return nil
}

func (cashAppReport *Report) GetTransactions() []Transaction {
	return cashAppReport.transactions
}

func NewCashAppReport(fileRecordsGetter file.RecordsGetter) Report {
	return Report{fileRecordsGetter: fileRecordsGetter, transactions: []Transaction{}}
}
