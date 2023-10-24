package ynab_test

import (
	"testing"

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
}
