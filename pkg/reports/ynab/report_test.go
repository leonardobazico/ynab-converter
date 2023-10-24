package ynab_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"cash2ynab/pkg/reports"
	"cash2ynab/pkg/reports/ynab"
)

func TestYnabRecordTransformer(t *testing.T) {
	t.Parallel()

	t.Run("should implement report.ReportExporter interface", func(t *testing.T) {
		t.Parallel()

		// Given
		ynabRecordTransformer := ynab.NewYnabRecordTransformer()
		// When
		_, implementsInterface := interface{}(&ynabRecordTransformer).(reports.TransactionToRecordTransformer)
		// Then
		assert.True(t, implementsInterface, "YnabRecordTransformer does not implement ReportExporter interface")
	})

	t.Run("should return the header", func(t *testing.T) {
		t.Parallel()

		// Given
		ynabRecordTransformer := ynab.NewYnabRecordTransformer()
		// When
		header := ynabRecordTransformer.GetHeader()
		// Then
		assert.Equal(t, []string{"Date", "Payee", "Memo", "Amount"}, header)
	})

	t.Run("should return a record", func(t *testing.T) {
		t.Parallel()

		// Given
		ynabRecordTransformer := ynab.NewYnabRecordTransformer()
		// When
		records, err := ynabRecordTransformer.GetRecords(
			[]reports.Transactioner{
				reports.Transaction{
					Counterparty: "MTA*NYCT PAYGO",
					Description:  "CARD CHARGED",
					Amount:       -2.9,
					Datetime:     time.Date(2023, 1, 2, 23, 59, 59, 0, time.UTC),
				},
				reports.Transaction{
					Counterparty: "Some business name",
					Description:  "PAYMENT RECEIVED",
					Amount:       10.0,
					Datetime:     time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC),
				},
			},
		)
		// Then
		assert.NoError(t, err)
		assert.Equal(
			t,
			[][]string{
				{
					"01/02/2023",
					"MTA*NYCT PAYGO",
					"CARD CHARGED",
					"-2.90",
				},
				{
					"12/31/2022",
					"Some business name",
					"PAYMENT RECEIVED",
					"10.00",
				},
			},
			records,
		)
	})
}
