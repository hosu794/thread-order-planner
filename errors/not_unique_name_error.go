package errors

type NotUniqueNameError struct {
	message string
}

func (e NotUniqueNameError) Error() string {
	return e.message
}

func NewNotUniqueNameError(name string) error {
	return NotUniqueNameError{"Given name " + name + " isn't unique"}
}
