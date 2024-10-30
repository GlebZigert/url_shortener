package pkg1

import (
	"fmt"
	"os"
)

func mulfunc(i int) (int, error) {
	return i * 2, nil
}

// errCheckFunc is func for tests
func errCheckFunc() {
	// формулируем ожидания: анализатор должен находить ошибку,
	// описанную в комментарии want
	mulfunc(5)           // want "expression returns unchecked error"
	res, _ := mulfunc(5) // want "assignment with unchecked error"
	fmt.Println(res)     // want "expression returns unchecked error"
	go mulfunc(5)
	defer mulfunc(5)
}

// FF is func for tests
func FF() {
	os.Remove(`myfile.txt`)
	os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
}
