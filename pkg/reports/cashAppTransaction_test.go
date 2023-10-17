package reports_test

import (
	"testing"

	"cash2ynab/pkg/reports"
)

func TestCashAppTransaction(t *testing.T) {
	t.Parallel()

	t.Run("should create a CashAppTransaction from a record", func(t *testing.T) {
		t.Parallel()
		// Given
		record := []string{
			"rmgsrz",
			"2023-10-06 19:56:39 EDT",
			"Cash Card Debit",
			"USD",
			"-$2.90",
			"$0",
			"-$2.90",
			"",
			"",
			"",
			"CARD CHARGED",
			"MTA*NYCT PAYGO",
			"",
			"Visa Debit 0987",
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
		// When
		cashAppTransaction := reports.NewCashAppTransaction(record)
		// Then
		if cashAppTransaction != expectedCashAppTransaction {
			t.Errorf("Expected cashAppTransaction to be %v. Got %v", expectedCashAppTransaction, cashAppTransaction)
		}
	})

	t.Run("should implement Transaction interface", func(t *testing.T) {
		t.Parallel()

		// Given
		record := []string{
			"rmgsrz",
			"2023-10-06 19:56:39 EDT",
			"Cash Card Debit",
			"USD",
			"-$2.90",
			"$0",
			"-$2.90",
			"",
			"",
			"",
			"CARD CHARGED",
			"MTA*NYCT PAYGO",
			"",
			"Visa Debit 0987",
		}
		// When
		cashAppTransaction := reports.NewCashAppTransaction(record)
		// Then

		t.Run("should GetCounterparty", func(t *testing.T) {
			if cashAppTransaction.GetCounterparty() != "MTA*NYCT PAYGO" {
				t.Errorf(
					"Expected Counterparty to be MTA*NYCT PAYGO. Got %v",
					cashAppTransaction.GetCounterparty(),
				)
			}
		})
	})
}
