package file

type FileRecordsGetter interface {
	GetRecordsFrom(filePath string) ([][]string, error)
}
