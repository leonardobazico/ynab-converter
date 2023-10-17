package reports_test

import (
	"testing"

	"cash2ynab/internal/file"
	"cash2ynab/pkg/reports"
	utils_test "cash2ynab/tests/utils"
)

// CashAppReport is a struct that represents a Cash App report
// it should have transactions represented by an array of CashAppTransaction
// CashAppReport should be able to transform records from a CSV file into an array of CashAppTransaction

func TestCashAppReport(t *testing.T) {
	t.Parallel()

	t.Run("should parse file and get an array of CashAppTransaction", func(t *testing.T) {
		// Given
		cashAppReport := reports.NewCashAppReport(file.NewCsvReader(utils_test.ExampleFilesFS))
		// When
		err := cashAppReport.ParseFileRecords("examples/cash_app_report_one_transaction.csv")
		transactions := cashAppReport.GetTransactions()
		// Then
		if err != nil {
			t.Errorf("Expected no error. Got %v", err)
		}
		if len(transactions) != 1 {
			t.Errorf("Expected transactions to have length 1. Got %v", len(transactions))
		}
		expectedCashAppTransaction := reports.CashAppTransaction{
			TransactionID:        "rmgsrz",
			Date:                 "2023-10-06 19:56:39 EDT",
			TransactionType:      "Cash Card Debit",
			Currency:             "USD",
			Amount:               "-$2.90",
			Fee:                  "$0",
			NetAmount:            "-$2.90",
			AssetType:            "",
			AssetPrice:           "",
			AssetAmount:          "",
			Status:               "CARD CHARGED",
			Notes:                "MTA*NYCT PAYGO",
			NameOfSenderReceiver: "",
			Account:              "Visa Debit 0987",
		}
		if transactions[0] != expectedCashAppTransaction {
			t.Errorf("Expected transactions[0] to be %v. Got %v", expectedCashAppTransaction, transactions[0])
		}
	})
}
