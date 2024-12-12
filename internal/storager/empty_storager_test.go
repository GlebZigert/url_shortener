package storager

import (
	"testing"
)

func TestEmptyStoreLoad(t *testing.T) {

	store, err := NewEmptyStorager()

	if err != nil {
		t.Error(err.Error())
	}
	tests := []struct {
		name string
		len  int
	}{
		{
			name: "1",
			len:  0,
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

func TestEmptyStoreWrite(t *testing.T) {

	store, err := NewEmptyStorager()

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

func TestEmptyStoragerDelete(t *testing.T) {

	store, err := NewEmptyStorager()

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
