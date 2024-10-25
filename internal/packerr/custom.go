package packerr

type ErrConflict409 struct {
	S string
}

type ErrDeleted struct {
	S string
}

func (e *ErrDeleted) Error() string {
	return e.S
}

func (e *ErrConflict409) Error() string {
	return e.S
}
