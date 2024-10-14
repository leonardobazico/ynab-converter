package cashapp

import (
	"fmt"

	"ynabconverter/internal/file"
	"ynabconverter/pkg/reports"
)

type Importer struct {
	fileRecordsGetter file.RecordsGetter
	transactions      []reports.Transactioner
}

func (cashApp *Importer) ParseFileRecords(filePath string) error {
	records, err := cashApp.fileRecordsGetter.GetRecordsFrom(filePath)
	if err != nil {
		return fmt.Errorf("fail to get records from file: %w", err)
	}

	for _, record := range records {
		cashAppTransaction := NewCashAppTransaction(record)
		cashApp.transactions = append(
			cashApp.transactions,
			&cashAppTransaction,
		)
	}

	return nil
}

func (cashApp *Importer) GetTransactions() []reports.Transactioner {
	return cashApp.transactions
}

func NewCashAppReport(fileRecordsGetter file.RecordsGetter) Importer {
	return Importer{fileRecordsGetter: fileRecordsGetter, transactions: []reports.Transactioner{}}
}

func NewCashAppReportCsvImporter() Importer {
	return NewCashAppReport(file.NewCsvImporter())
}
