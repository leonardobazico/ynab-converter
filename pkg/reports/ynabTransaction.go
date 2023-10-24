package reports

import (
	"fmt"
)

type YnabTransaction struct {
	Date   string
	Payee  string
	Memo   string
	Amount string
}

func NewYnabTransaction(transaction Transactioner) (*YnabTransaction, error) {
	datetime, err := transaction.GetDatetime()
	if err != nil {
		return nil, fmt.Errorf("error getting datetime: %w", err)
	}

	amount, err := transaction.GetAmount()
	if err != nil {
		return nil, fmt.Errorf("error getting amount: %w", err)
	}

	return &YnabTransaction{
		Date:   datetime.Format("01/02/2006"),
		Payee:  transaction.GetCounterparty(),
		Memo:   transaction.GetDescription(),
		Amount: fmt.Sprintf("%.2f", amount),
	}, nil
}
