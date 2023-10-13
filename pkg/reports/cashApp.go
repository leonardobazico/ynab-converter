package reports

import (
	"cash2ynab/internal/file"
)

type CashAppReport struct {
	fileRecordsGetter file.FileRecordsGetter
	transactions      []CashAppTransaction
}

func (report *CashAppReport) ParseFileRecords(filePath string) error {
	records, err := report.fileRecordsGetter.GetRecordsFrom(filePath)
	if err != nil {
		return err
	}
	for _, record := range records {
		report.transactions = append(report.transactions, NewCashAppTransaction(record))
	}
	return nil
}

func (report *CashAppReport) GetTransactions() []CashAppTransaction {
	return report.transactions
}

func NewCashAppReport(fileRecordsGetter file.FileRecordsGetter) CashAppReport {
	return CashAppReport{fileRecordsGetter: fileRecordsGetter, transactions: []CashAppTransaction{}}
}
