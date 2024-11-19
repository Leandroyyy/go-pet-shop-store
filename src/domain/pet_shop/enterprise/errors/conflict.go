package enterprise_errors

type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return e.Message
}

func NewConflictError(message string) error {
	return &ConflictError{Message: message}
}
