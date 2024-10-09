package cashapp

import (
	"fmt"

	"ynabconverter/internal/file"
	"ynabconverter/pkg/reports"
)

type ReportImporter struct {
	fileRecordsGetter file.RecordsGetter
	transactions      []reports.Transactioner
}

func (cashApp *ReportImporter) ParseFileRecords(filePath string) error {
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

func (cashApp *ReportImporter) GetTransactions() []reports.Transactioner {
	return cashApp.transactions
}

func NewCashAppReport(fileRecordsGetter file.RecordsGetter) ReportImporter {
	return ReportImporter{fileRecordsGetter: fileRecordsGetter, transactions: []reports.Transactioner{}}
}

func NewCashAppReportCsvImporter() ReportImporter {
	return NewCashAppReport(file.NewCsvImporter())
}
