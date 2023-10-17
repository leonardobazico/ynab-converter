package reports

import (
	"fmt"

	"cash2ynab/internal/file"
)

type CashAppReport struct {
	fileRecordsGetter file.RecordsGetter
	transactions      []CashAppTransaction
}

func (report *CashAppReport) ParseFileRecords(filePath string) error {
	records, err := report.fileRecordsGetter.GetRecordsFrom(filePath)
	if err != nil {
		return fmt.Errorf("fail to get records from file: %w", err)
	}

	for _, record := range records {
		report.transactions = append(report.transactions, NewCashAppTransaction(record))
	}

	return nil
}

func (report *CashAppReport) GetTransactions() []CashAppTransaction {
	return report.transactions
}

func NewCashAppReport(fileRecordsGetter file.RecordsGetter) CashAppReport {
	return CashAppReport{fileRecordsGetter: fileRecordsGetter, transactions: []CashAppTransaction{}}
}
