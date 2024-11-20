package filereader

import (
	"bufio"
	"fmt"
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

// GetReader return reader
func (GetFileReader) Write(data []byte, filepath string) error {

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer packerr.AddCloseErrToErr(&err, file)
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	writer.Flush()

	return err

}

func (GetFileReader) CheckFile(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return err
}

// New is constructor
func New() *GetFileReader {
	return &GetFileReader{}
}
