package file_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"cash2ynab/internal/file"
	utils_test "cash2ynab/tests/utils"
)

func TestReadCsv(t *testing.T) {
	t.Parallel()

	t.Run("should return error when file does not exist", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "examples/does-not-exist.csv"
		csvReader := file.NewCsvReader(utils_test.ExampleFilesFS)
		// When
		_, err := csvReader.GetRecordsFrom(filePath)
		// Then
		if err == nil {
			t.Errorf("Expected an error when the file does not exist")
		}
		if err.Error() != "fail to open file: open examples/does-not-exist.csv: file does not exist" {
			t.Errorf("Expected a different error message. Got:\n%v", err.Error())
		}
	})

	t.Run("should return error when file exists but is not a csv", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "examples/not-a-csv.txt"
		csvReader := file.NewCsvReader(utils_test.ExampleFilesFS)
		// When
		output, err := csvReader.GetRecordsFrom(filePath)
		// Then
		assert.Nil(t, output)
		assert.ErrorContains(t, err, "fail to read csv file:")
	})

	t.Run("should not return error when file exists", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "examples/cash_app_report.csv"
		csvReader := file.NewCsvReader(utils_test.ExampleFilesFS)
		// When
		_, err := csvReader.GetRecordsFrom(filePath)
		// Then
		if err != nil {
			t.Errorf("Expected no error when the file exists:\n%v", err)
		}
	})

	t.Run("should not return error when file exists and is empty", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "examples/empty.csv"
		csvReader := file.NewCsvReader(utils_test.ExampleFilesFS)
		// When
		records, err := csvReader.GetRecordsFrom(filePath)
		// Then
		if err != nil {
			t.Errorf("Expected no error when the file exists and is empty")
		}
		if len(records) > 0 {
			t.Errorf("Expected records to be empty")
		}
	})

	t.Run("should ignore title from records", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "examples/just_title.csv"
		csvReader := file.NewCsvReader(utils_test.ExampleFilesFS)
		// When
		records, err := csvReader.GetRecordsFrom(filePath)
		if err != nil {
			t.Errorf("Expected no error when the file exists and is empty")
		}
		// Then
		if len(records) > 0 {
			t.Errorf("Expected records to be empty")
		}
	})

	t.Run("should return matrix of strings when file exists and is not empty", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "examples/cash_app_report.csv"
		csvReader := file.NewCsvReader(utils_test.ExampleFilesFS)
		// When
		records, err := csvReader.GetRecordsFrom(filePath)
		// Then
		if err != nil {
			t.Errorf("Expected no error when the file exists and is not empty")
		}
		if len(records) == 0 {
			t.Errorf("Expected records to have at least one row")
		}
	})
}
