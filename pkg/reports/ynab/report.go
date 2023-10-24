package ynab

type TransactionToRecordTransformer struct {
	header []string
}

func (ynab TransactionToRecordTransformer) GetHeader() []string {
	return ynab.header
}

func NewYnabRecordTransformer() TransactionToRecordTransformer {
	return TransactionToRecordTransformer{
		header: []string{"Date", "Payee", "Memo", "Amount"},
	}
}
