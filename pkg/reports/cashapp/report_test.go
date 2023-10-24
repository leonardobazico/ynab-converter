package cashapp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"cash2ynab/internal/file"
	"cash2ynab/pkg/reports/cashapp"
	utils_test "cash2ynab/tests/utils"
)

func TestCashAppReport(t *testing.T) {
	t.Parallel()

	t.Run("should return error when file does not exist", func(t *testing.T) {
		t.Parallel()

		// Given
		cashAppReport := cashapp.NewCashAppReport(file.NewCsvReader(utils_test.ExampleFilesFS))
		// When
		err := cashAppReport.ParseFileRecords("examples/does-not-exist.csv")
		// Then
		assert.Error(t, err)
	})

	t.Run("should parse file and get an array of cashapp.Transaction", func(t *testing.T) {
		t.Parallel()

		// Given
		cashAppReport := cashapp.NewCashAppReport(file.NewCsvReader(utils_test.ExampleFilesFS))
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
		expectedCashAppTransaction := cashapp.Transaction{
			TransactionID:        "rmgsrz",
			Date:                 "2023-10-06 23:59:59 EDT",
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
