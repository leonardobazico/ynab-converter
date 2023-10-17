package file

type RecordsGetter interface {
	GetRecordsFrom(filePath string) ([][]string, error)
}
