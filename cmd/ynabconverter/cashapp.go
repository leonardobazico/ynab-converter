package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"ynabconverter/pkg/reports/cashapp"
	"ynabconverter/pkg/reports/ynab"
)

func cashApp(cashAppCmd *flag.FlagSet, cashAppFile *string) {
	if cashAppCmd.Parse(os.Args[2:]) != nil {
		os.Exit(1)
	}

	if *cashAppFile == "" {
		fmt.Fprintf(os.Stderr, "it is required to set a file to be converted\n\n")
		cashAppCmd.Usage()
		os.Exit(1)
	}

	cashAppImporter := cashapp.NewCashAppReportCsvImporter()
	err := cashAppImporter.ParseFileRecords(*cashAppFile)
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
