package filereader

import (
	"testing"
)

func TestReadFile(t *testing.T) {

	test := struct {
		name string
	}{
		name: "filereader test",
	}
	t.Run(test.name, func(t *testing.T) {
		fr := New()
		fr.Read("./test.txt")

	})
}

func TestFailReadFile(t *testing.T) {

	test := struct {
		name string
	}{
		name: "filereader test",
	}
	t.Run(test.name, func(t *testing.T) {
		fr := New()
		fr.Read("./no_test.txt")

	})
}

func TestWriteFile(t *testing.T) {

	test := struct {
		name string
	}{
		name: "filereader test",
	}
	t.Run(test.name, func(t *testing.T) {
		fr := New()
		bytes := []byte("qwerty")
		fr.Write(bytes, "./test.txt")

	})
}

func TestXheckFile(t *testing.T) {

	test := struct {
		name string
	}{
		name: "filereader test",
	}
	t.Run(test.name, func(t *testing.T) {
		fr := New()

		fr.CheckFile("./test.txt")

	})
}
func TestCheckFile(t *testing.T) {

	test := struct {
		name string
	}{
		name: "filereader test",
	}
	t.Run(test.name, func(t *testing.T) {
		fr := New()

		fr.CheckFile("./test1.txt")

	})
}

func TestNoCheckFile(t *testing.T) {

	test := struct {
		name string
	}{
		name: "filereader test",
	}
	t.Run(test.name, func(t *testing.T) {
		fr := New()

		fr.CheckFile("./no_test.txt")

	})
}
