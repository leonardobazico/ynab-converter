package file

import (
	"encoding/csv"
	"io/fs"
	"log"
)

type csvReader struct {
	fileSystem fs.FS
}

func (reader *csvReader) GetRecordsFrom(filePath string) ([][]string, error) {
	csvFile, err := reader.fileSystem.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	ignoreRecord(csvReader)

	remainingRecords, err := csvReader.ReadAll()
	return remainingRecords, err
}

func ignoreRecord(csvReader *csv.Reader) {
	_, err := csvReader.Read()
	if err != nil {
		log.Default().Println("Error ignoring record", err)
	}
}

func NewCsvReader(fs fs.FS) FileRecordsGetter {
	return &csvReader{fileSystem: fs}
}
