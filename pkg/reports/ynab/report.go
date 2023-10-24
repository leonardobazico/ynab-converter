package ynab

import (
	"fmt"

	"cash2ynab/pkg/reports"
)

type TransactionToRecordTransformer struct {
	header []string
}

func (ynab TransactionToRecordTransformer) GetHeader() []string {
	return ynab.header
}

func (ynab TransactionToRecordTransformer) GetRecords(transactions []reports.Transactioner) ([][]string, error) {
	records := [][]string{}
	for _, transaction := range transactions {
		ynabTransaction, err := NewYnabTransaction(transaction)
		if err != nil {
			return nil, fmt.Errorf("error creating ynab transaction: %w", err)
		}

		records = append(records, []string{
			ynabTransaction.Date,
			ynabTransaction.Payee,
			ynabTransaction.Memo,
			ynabTransaction.Amount,
		})
	}

	return records, nil
}

func NewYnabRecordTransformer() TransactionToRecordTransformer {
	return TransactionToRecordTransformer{
		header: []string{"Date", "Payee", "Memo", "Amount"},
	}
}
