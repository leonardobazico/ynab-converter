package ynab

import (
	"fmt"

	"ynabconverter/pkg/reports"
)

type Transaction struct {
	Date   string
	Payee  string
	Memo   string
	Amount string
}

func NewYnabTransaction(transaction reports.Transactioner) (*Transaction, error) {
	datetime, err := transaction.GetDatetime()
	if err != nil {
		return nil, fmt.Errorf("error getting datetime: %w", err)
	}

	amount, err := transaction.GetAmount()
	if err != nil {
		return nil, fmt.Errorf("error getting amount: %w", err)
	}

	return &Transaction{
		Date:   datetime.Format("01/02/2006"),
		Payee:  transaction.GetCounterparty(),
		Memo:   transaction.GetDescription(),
		Amount: fmt.Sprintf("%.2f", amount),
	}, nil
}
