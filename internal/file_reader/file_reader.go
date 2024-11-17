package file_reader

import (
	"bufio"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

type GetFileReader struct {
}

func (GetFileReader) GetReader(path string) (*bufio.Reader, error) {

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer packerr.AddCloseErrToErr(&err, file)

	return bufio.NewReader(file), nil
}

func New() *GetFileReader {
	return &GetFileReader{}
}
