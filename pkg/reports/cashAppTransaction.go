package reports

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CashAppTransaction struct {
	TransactionID        string
	Date                 string
	TransactionType      string
	Currency             string
	Amount               string
	Fee                  string
	NetAmount            string
	AssetType            string
	AssetPrice           string
	AssetAmount          string
	Status               string
	Notes                string
	NameOfSenderReceiver string
	Account              string
}

func (cashAppTransaction *CashAppTransaction) GetCounterparty() string {
	return cashAppTransaction.Notes
}

func (cashAppTransaction *CashAppTransaction) GetDescription() string {
	return cashAppTransaction.Status
}

func (cashAppTransaction *CashAppTransaction) GetAmount() (float32, error) {
	amountWithoutDollarSign := strings.ReplaceAll(cashAppTransaction.Amount, "$", "")

	amount, err := strconv.ParseFloat(amountWithoutDollarSign, 32)
	if err != nil {
		return 0, fmt.Errorf("error parsing amount: %w", err)
	}

	return float32(amount), nil
}

func (cashAppTransaction *CashAppTransaction) GetDatetime() (*time.Time, error) {
	datetime, err := cashAppTransaction.dateToDatetime()
	if err != nil {
		return nil, fmt.Errorf("error getting datetime: %w", err)
	}

	return datetime, nil
}

func NewCashAppTransaction(record []string) CashAppTransaction {
	return CashAppTransaction{
		TransactionID:        record[0],
		Date:                 record[1],
		TransactionType:      record[2],
		Currency:             record[3],
		Amount:               record[4],
		Fee:                  record[5],
		NetAmount:            record[6],
		AssetType:            record[7],
		AssetPrice:           record[8],
		AssetAmount:          record[9],
		Status:               record[10],
		Notes:                record[11],
		NameOfSenderReceiver: record[12],
		Account:              record[13],
	}
}

func (cashAppTransaction *CashAppTransaction) dateToDatetime() (*time.Time, error) {
	easternDaylightTime := "EDT"
	dateNormalized := strings.ReplaceAll(cashAppTransaction.Date, easternDaylightTime, "EST")
	locationString, err := getLocationString(dateNormalized)
	if err != nil {
		return nil, fmt.Errorf("error getting location string: %w", err)
	}

	location, _ := time.LoadLocation(locationString)
	datetime, _ := time.ParseInLocation(time.DateTime+" MST", dateNormalized, location)

	isEasternDaylightTime := strings.Contains(cashAppTransaction.Date, easternDaylightTime)
	if isEasternDaylightTime {
		fixedEdtToEst := datetime.Add(-time.Hour)

		return &fixedEdtToEst, nil
	}

	return &datetime, nil
}

func getLocationString(transactionDate string) (string, error) {
	datetimeToExtractLocation, err := time.Parse(time.DateTime+" MST", transactionDate)
	if err != nil {
		return "", fmt.Errorf("error parsing datetime: %w", err)
	}

	timezoneAbbreviation := datetimeToExtractLocation.Location().String()

	return timezoneAbbreviation, nil
}
