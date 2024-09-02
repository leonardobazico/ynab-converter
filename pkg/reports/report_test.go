package reports_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"cash2ynab/pkg/reports"
)

func TestTransaction(t *testing.T) {
	t.Parallel()

	t.Run("should implement report.Transactioner interface", func(t *testing.T) {
		t.Parallel()

		// Given
		transaction := reports.Transaction{
			Counterparty: "MTA*NYCT PAYGO",
			Description:  "CARD CHARGED",
			Amount:       -2.9,
			Datetime:     time.Date(2023, 10, 6, 23, 59, 59, 0, time.UTC),
		}
		// When
		_, implementsInterface := interface{}(&transaction).(reports.Transactioner)
		// Then

		assert.True(t, implementsInterface, "Transaction does not implement Transactioner interface")

		t.Run("should return the counterparty", func(t *testing.T) {
			// When
			counterparty := transaction.GetCounterparty()
			// Then
			assert.Equal(t, "MTA*NYCT PAYGO", counterparty)
		})

		t.Run("should return the description", func(t *testing.T) {
			// When
			description := transaction.GetDescription()
			// Then
			assert.Equal(t, "CARD CHARGED", description)
		})

		t.Run("should return the amount", func(t *testing.T) {
			// When
			amount, _ := transaction.GetAmount()
			// Then
			assert.InEpsilon(t, float32(-2.9), amount, 0.0001)
		})

		t.Run("should return the datetime", func(t *testing.T) {
			// When
			datetime, _ := transaction.GetDatetime()
			// Then
			assert.Equal(t, time.Date(2023, 10, 6, 23, 59, 59, 0, time.UTC), *datetime)
		})
	})
}
