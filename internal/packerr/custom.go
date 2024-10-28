package packerr

// ошибка когда создаем шорт но он уже есть
type ErrConflict409 struct {
	S string
}

var Conflict ErrConflict409

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
