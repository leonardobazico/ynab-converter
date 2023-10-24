package reports_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"cash2ynab/pkg/reports"
)

//nolint:paralleltest
func TestCashAppTransaction(t *testing.T) {
	t.Run("should create a CashAppTransaction from a record", func(t *testing.T) {
		// Given
		record := []string{
			"rmgsrz",
			"2023-10-06 23:59:59 EDT",
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
		// When
		cashAppTransaction := reports.NewCashAppTransaction(record)
		// Then
		if cashAppTransaction != expectedCashAppTransaction {
			t.Errorf("Expected cashAppTransaction to be %v. Got %v", expectedCashAppTransaction, cashAppTransaction)
		}
	})

	t.Run("should implement report.Transaction interface", func(t *testing.T) {
		// Given
		record := []string{
			"rmgsrz",
			"2023-10-06 23:59:59 EDT",
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
		_, implementsInterface := interface{}(&cashAppTransaction).(reports.Transaction)
		assert.True(t, implementsInterface, "CashAppTransaction does not implement Transaction interface")

		t.Run("should GetCounterparty", func(t *testing.T) {
			assert.Equal(t, "MTA*NYCT PAYGO", cashAppTransaction.GetCounterparty())
		})

		t.Run("should GetDescription", func(t *testing.T) {
			assert.Equal(t, "CARD CHARGED", cashAppTransaction.GetDescription())
		})

		t.Run("should return error when amount is not a float", func(t *testing.T) {
			// Given
			cashAppNotFloatTransaction := reports.CashAppTransaction{
				Amount: "not a float",
			}
			// When
			_, err := cashAppNotFloatTransaction.GetAmount()
			// Then
			assert.Error(t, err)
		})

		t.Run("should GetAmount", func(t *testing.T) {
			amount, _ := cashAppTransaction.GetAmount()
			assert.Equal(t, float32(-2.90), amount)
		})

		t.Run("should return error when datetime is not valid", func(t *testing.T) {
			// Given
			cashAppNotValidDateTransaction := reports.CashAppTransaction{
				Date: "not a valid date",
			}
			// When
			_, err := cashAppNotValidDateTransaction.GetDatetime()
			// Then
			assert.Error(t, err)
		})

		t.Run("should GetDatetime", func(t *testing.T) {
			// Given
			t.Setenv("TZ", "UTC")
			cashAppTransaction := reports.CashAppTransaction{
				Date: "2023-12-06 23:59:59 EST",
			}
			// When
			datetime, _ := cashAppTransaction.GetDatetime()
			// Then
			estLocation, _ := time.LoadLocation("EST")
			expectedDatetime := time.Date(2023, 12, 6, 23, 59, 59, 0, estLocation)
			assert.Equal(t, &expectedDatetime, datetime)
			assert.Equal(t, time.Date(2023, 12, 7, 4, 59, 59, 0, time.UTC), datetime.UTC())
		})

		t.Run("should get Eastern Daylight Time as EST offset by 1h", func(t *testing.T) {
			// Given
			t.Setenv("TZ", "UTC")
			cashAppTransaction := reports.CashAppTransaction{
				Date: "2023-10-06 23:59:59 EDT",
			}
			// When
			datetime, _ := cashAppTransaction.GetDatetime()
			// Then
			estLocation, _ := time.LoadLocation("EST")
			expectedDatetime := time.Date(2023, 10, 6, 22, 59, 59, 0, estLocation)
			assert.Equal(t, &expectedDatetime, datetime)
			assert.Equal(t, time.Date(2023, 10, 7, 3, 59, 59, 0, time.UTC), datetime.UTC())
		})
	})
}
