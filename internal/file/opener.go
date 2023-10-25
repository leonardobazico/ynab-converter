package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Opener interface {
	Open(filePath string) (fs.File, error)
}

type OpenWrapper struct {
	fileSystem fs.FS
}

//nolint:wrapcheck
func (opener *OpenWrapper) Open(filePath string) (fs.File, error) {
	cleanedFilePath := filepath.Clean(filePath)

	if opener.fileSystem == nil {
		var readOnlyFileMode fs.FileMode = 0400

		return os.OpenFile(cleanedFilePath, os.O_RDONLY, readOnlyFileMode)
	}

	return opener.fileSystem.Open(cleanedFilePath)
}

func NewFileSytemOpener(fs fs.FS) *OpenWrapper {
	return &OpenWrapper{fileSystem: fs}
}

func NewOsOpener() *OpenWrapper {
	return &OpenWrapper{fileSystem: nil}
}
