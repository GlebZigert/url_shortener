package storager

import (
	"testing"

	"github.com/GlebZigert/url_shortener.git/storemocks"
	"github.com/golang/mock/gomock"
)

func TestFStoreLoad(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	mockreader := storemocks.NewMockFileReaderWriter(ctrl)

	mockreader.EXPECT().Read(gomock.Any()).DoAndReturn(func(path string) ([]byte, error) {
		/*
			data := []byte(`{"uuid":"1","short_url":"4rSPg8ap","original_url":"http://yandex.ru"}
			{"uuid":"2","short_url":"edVPg3ks","original_url":"http://ya.ru"}
			{"uuid":"3","short_url":"dG56Hqxm","original_url":"http://practicum.yandex.ru"}
				`)
		*/

		data := []byte(`{"uuid":"1","short_url":"4rSPg8ap","original_url":"http://yandex.ru"}
					`)

		return data, nil
	}).Times(1)

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
			store.Load(&shorten)

			t.Log(len(shorten))

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
			name: "1",
			len:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			store.StorageWrite("11", "222", 0)

		})
	}

}
