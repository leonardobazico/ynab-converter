package main

import (
	"encoding/csv"
	"log"
	"os"

	"ynabconverter/pkg/reports/cashapp"
	"ynabconverter/pkg/reports/ynab"
)

func main() {
	cashAppFile := os.Args[1]

	cashAppImporter := cashapp.NewCashAppReportCsvImporter()
	err := cashAppImporter.ParseFileRecords(cashAppFile)
	if err != nil {
		log.Fatalf("Error parsing file records: %v", err)
	}

	transactions := cashAppImporter.GetTransactions()

	ynabRecordTransformer := ynab.NewYnabRecordTransformer()
	records, err := ynabRecordTransformer.GetRecordsWithHeader(transactions)
	if err != nil {
		log.Fatalf("Error getting YNAB records: %v", err)
	}

	writer := csv.NewWriter(os.Stdout)
	err = writer.WriteAll(records)
	if err != nil {
		log.Fatalf("Error writing records: %v", err)
	}
}
