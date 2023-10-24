package file

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"log"
)

type CsvImporter struct {
	fileSystem fs.FS
}

func (reader *CsvImporter) GetRecordsFrom(filePath string) ([][]string, error) {
	csvFile, err := reader.fileSystem.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open file: %w", err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	ignoreRecord(csvReader)

	remainingRecords, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("fail to read csv file: %w", err)
	}

	return remainingRecords, nil
}

func ignoreRecord(csvReader *csv.Reader) {
	_, err := csvReader.Read()
	if err != nil {
		log.Default().Println("Error ignoring record", err)
	}
}

func NewCsvImporter(fs fs.FS) *CsvImporter {
	return &CsvImporter{fileSystem: fs}
}
