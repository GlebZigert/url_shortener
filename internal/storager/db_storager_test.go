package storager

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/storemocks"
	"github.com/golang/mock/gomock"
)

func TestDBStoragerLoad(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	tdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer tdb.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("SELECT * FROM strazh").WillReturnRows(rows)

	store, err := NewDBStorager(cfg, db.Get(tdb))

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

func TestDBStoragerInsert(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	tdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer tdb.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("SELECT * FROM strazh").WillReturnRows(rows)

	store, err := NewDBStorager(cfg, db.Get(tdb))
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
			store.StorageWrite("", "", 1)

			t.Log(len(shorten))

		})
	}

}

func TestDBStoragerDelete(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	tdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer tdb.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("SELECT * FROM strazh").WillReturnRows(rows)

	store, err := NewDBStorager(cfg, db.Get(tdb))
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
			store.Delete("1")

			t.Log(len(shorten))

		})
	}

}

func TestDBStoragerDeleteList(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	tdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer tdb.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("SELECT * FROM strazh").WillReturnRows(rows)

	store, err := NewDBStorager(cfg, db.Get(tdb))
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
			store.Delete([]string{"1", "2"})
			store.Delete(1)
			t.Log(len(shorten))

		})
	}

}

func TestDBStoragerLoad1(t *testing.T) {

	ctrl := gomock.NewController(t)
	cfg := storemocks.NewMockStoreConfig(ctrl)

	tdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer tdb.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("SELECT * FROM strazh").WillReturnRows(rows)

	store := New(cfg, filereader.New(), db.Get(tdb))

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
