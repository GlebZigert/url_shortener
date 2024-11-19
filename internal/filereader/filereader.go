package filereader

import (
	"bufio"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

// GetFileReader to read from file
type GetFileReader struct {
}

// GetReader return reader
func (GetFileReader) Read(path string) ([]byte, error) {

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}
	defer packerr.AddCloseErrToErr(&err, file)

	reader := bufio.NewReader(file)
	var data []byte
	_, err = reader.Read(data)

	return data, err
}

// New is constructor
func New() *GetFileReader {
	return &GetFileReader{}
}
