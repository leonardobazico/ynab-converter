package cashapp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"cash2ynab/internal/file"
	"cash2ynab/pkg/reports"
	"cash2ynab/pkg/reports/cashapp"
	utils_test "cash2ynab/tests/utils"
)

func TestCashAppReportImporter(t *testing.T) {
	t.Parallel()

	t.Run("should return error when file does not exist", func(t *testing.T) {
		t.Parallel()

		// Given
		cashAppReport := cashapp.NewCashAppReport(file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS))
		// When
		err := cashAppReport.ParseFileRecords("examples/does-not-exist.csv")
		// Then
		assert.Error(t, err)
	})

	t.Run("should parse file and get an array of cashapp.Transaction", func(t *testing.T) {
		t.Parallel()

		// Given
		cashAppReport := cashapp.NewCashAppReport(file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS))
		// When
		err := cashAppReport.ParseFileRecords("examples/cash_app_report_sample.csv")
		transactions := cashAppReport.GetTransactions()
		// Then
		assert.NoError(t, err)
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
		assert.Equal(t, &expectedCashAppTransaction, transactions[0])
	})

	t.Run("should implement report.ReportImporter interface", func(t *testing.T) {
		t.Parallel()

		// Given
		cashAppReport := cashapp.NewCashAppReport(file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS))
		// When
		_, implementsInterface := interface{}(&cashAppReport).(reports.ReportImporter)
		// Then
		assert.True(t, implementsInterface, "CashAppReport does not implement ReportImporter interface")
	})
}
