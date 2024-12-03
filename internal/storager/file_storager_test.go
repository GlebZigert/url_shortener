package storager

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/GlebZigert/url_shortener.git/storemocks"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestFStoreLoad(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	mockreader := storemocks.NewMockFileReaderWriter(ctrl)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		data := []byte(`{"uuid":1,"short_url":"4rSPg8ap","original_url":"http://yandex.ru"}
					`)
		return data, nil
	}).Times(1)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		data := []byte(`{"uid":1}`)
		return data, nil
	}).Times(1)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		data := []byte(`{"uid":2}`)
		return data, nil
	}).Times(1)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		data := []byte(`{"uuid":2,"short_url":"4rSPg8ap","original_url":"http://yandex.ru"}
					`)
		return data, nil
	}).Times(1)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		return nil, errors.New("")
	}).AnyTimes()
	mockreader.EXPECT().CheckFile(gomock.Any()).DoAndReturn(func(path string) error {
		return nil
	}).AnyTimes()

	cfg.EXPECT().GetFileStoragePath().DoAndReturn(func() string {
		return "any path"
	}).AnyTimes()

	cfg.EXPECT().GetDatabaseDSN().DoAndReturn(func() string {
		return "any dsn"
	}).AnyTimes()

	store, err := NewFileStorager(cfg, mockreader)

	if err != nil {
		t.Error(err.Error())
	}
	tests := []struct {
		name string
		len  int
	}{
		{
			name: "1",
			len:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			var shorten []*Shorten
			res, err := store.Load(&shorten)

			if err != nil {
				t.Error(err.Error())
			}

			t.Log(len(*res))

			assert.Equal(t, 2, len(*res))

		})
	}

}

func TestFStoreLoadUsers(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	mockreader := storemocks.NewMockFileReaderWriter(ctrl)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		data := []byte(`{"uuid":1,"short_url":"4rSPg8ap","original_url":"http://yandex.ru"}
					`)
		return data, nil
	}).Times(1)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		data := []byte(`{"uid":1}`)
		return data, nil
	}).Times(1)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		data := []byte(`{"uid":2}`)
		return data, nil
	}).Times(1)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		data := []byte(`{"uuid":2,"short_url":"4rSPg8ap","original_url":"http://yandex.ru"}
					`)
		return data, nil
	}).Times(1)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		return nil, errors.New("")
	}).AnyTimes()
	mockreader.EXPECT().CheckFile(gomock.Any()).DoAndReturn(func(path string) error {
		return nil
	}).AnyTimes()

	cfg.EXPECT().GetFileStoragePath().DoAndReturn(func() string {
		return "any path"
	}).AnyTimes()

	cfg.EXPECT().GetDatabaseDSN().DoAndReturn(func() string {
		return "any dsn"
	}).AnyTimes()

	store, err := NewFileStorager(cfg, mockreader)

	if err != nil {
		t.Error(err.Error())
	}
	tests := []struct {
		name string
		len  int
	}{
		{
			name: "1",
			len:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			var users []*User
			resusers, err := store.LoadUsers(&users)

			if err != nil {
				t.Error(err.Error())
			}
			assert.Equal(t, 2, len(*resusers))

		})
	}

}

func TestFStoreWrite(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	mockreader := storemocks.NewMockFileReaderWriter(ctrl)

	mockreader.EXPECT().Write(gomock.Any(), gomock.Any()).DoAndReturn(func(data []byte, filepath string) error {
		return nil
	}).AnyTimes()

	mockreader.EXPECT().CheckFile(gomock.Any()).DoAndReturn(func(path string) error {
		return nil
	}).AnyTimes()

	cfg.EXPECT().GetFileStoragePath().DoAndReturn(func() string {
		return "any path"
	}).AnyTimes()

	cfg.EXPECT().GetDatabaseDSN().DoAndReturn(func() string {
		return "any dsn"
	}).AnyTimes()

	store, err := NewFileStorager(cfg, mockreader)

	if err != nil {
		t.Error(err.Error())
	}
	tests := []struct {
		name string
		len  int
	}{
		{
			name: "01",
			len:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			err := store.StorageWrite("11", "222", 0)
			if err != nil {
				t.Error(err.Error())
			}

		})
	}

}

// MockFileWriter используется для имитации записи данных в файл.
type MockFileWriter struct {
	buffer bytes.Buffer //я создал его ради этого буффера.
	err    error
}

func (m *MockFileWriter) Write(p []byte, filepath string) error {
	if m.err != nil {
		return m.err
	}

	_, err := m.buffer.Write(p)

	return err
}

func (m *MockFileWriter) CheckFile(filepath string) error {
	return nil
}

func (m *MockFileWriter) Read(path string) ([]byte, error) {
	return []byte{}, nil
}

func TestFStoreWrite1(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	mockreader := &MockFileWriter{err: nil}

	cfg.EXPECT().GetFileStoragePath().DoAndReturn(func() string {
		return "any path"
	}).AnyTimes()

	cfg.EXPECT().GetDatabaseDSN().DoAndReturn(func() string {
		return "any dsn"
	}).AnyTimes()

	store, err := NewFileStorager(cfg, mockreader)

	if err != nil {
		t.Error(err.Error())
	}
	tests := []struct {
		name string
		len  int
	}{
		{
			name: "1",
			len:  1,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			t.Log("-->")

			err := store.StorageWrite("11", "222", 0)

			shorten := Shorten{1, 0, "11", "222", false}

			data, err := json.Marshal(&shorten)
			if err != nil {
				t.Errorf(err.Error())
			}

			if err != nil {
				t.Error(err.Error())
			}
			// буффер я создал чтобы получать эту строку - которую пишет в файл тестируемый метод

			res := mockreader.buffer.String()
			wanted := string(data) + "\n"
			t.Log("res   : ", res)
			t.Log("wanted: ", wanted)
			assert.Equal(t, strings.Compare(res, wanted), 0)

			t.Log("<--")
		})
	}

}
