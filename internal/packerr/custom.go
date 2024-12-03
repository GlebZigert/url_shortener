package packerr

// ошибка когда создаем шорт но он уже есть
type ErrConflict409 struct {
	S string
}

// обьявил ошибку
var Conflict ErrConflict409 = ErrConflict409{S: "попытка сократить уже имеющийся в базе URL"}

// ошибка когда просим шорт но он удален
type ErrDeleted struct {
	S string
}

// метод Error
func (e *ErrDeleted) Error() string {
	return e.S
}

// метод Error
func (e *ErrConflict409) Error() string {
	return e.S
}

// обьявил ошибку
var WrongType ErrWrongType = ErrWrongType{S: "неправильный тип"}

// ошибка когда просим шорт но он удален
type ErrWrongType struct {
	S string
}

// метод Error
func (e *ErrWrongType) Error() string {
	return e.S
}

// обьявил ошибку
var NewUserTryGetsURLs ErrWrongType = ErrWrongType{S: "новый юзер не создал ни одного url но сразу пытается сразу взять свои urlы "}

// ошибка когда просим шорт но он удален
type ErrNewUserTryGetsURLs struct {
	S string
}

// метод Error
func (e *ErrNewUserTryGetsURLs) Error() string {
	return e.S
}
