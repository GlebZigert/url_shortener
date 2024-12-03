package storager

import (
	"encoding/json"
	"testing"

	"github.com/GlebZigert/url_shortener.git/storemocks"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

// результатом выполнения функции WriteUID должна быть запись в файл
// "User: x" где x - следующий uid пользователя
// проконтролировать наличие и форму записи в буфер как результат выполнения функции WriteUID
func TestFStoreWriteUID(t *testing.T) {

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
		uid  int
	}{
		{
			name: "1",
			uid:  1,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			err = store.WriteUID(test.uid)

			if err != nil {
				t.Error(err.Error())
			}

			res := mockreader.buffer.String()
			t.Log("res   : ", res)

			var uid User

			err = json.Unmarshal(mockreader.buffer.Bytes(), &uid)

			if err != nil {
				t.Error(err.Error())
			}

			assert.Equal(t, uid.Uid, test.uid)

		})
	}

}
