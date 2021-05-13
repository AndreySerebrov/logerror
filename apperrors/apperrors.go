package apperrors

import "errors"

// Список ошибок уровня приложения
var (
	ErrNotFound        = errors.New("not found")
	ErrAccessForbidden = errors.New("access Forbidden")
	ErrInternalError   = errors.New("internal Error")
)

type ImportantError struct {
	err   error
	cause error
}

func NewImportantError(cause error) error {
	return ImportantError{err: errors.New("важная ошибка"), cause: cause}
}

func (err ImportantError) Error() string {
	return err.err.Error()
}

func (err *ImportantError) Unwrap() error {
	return err.cause
}

type VeryImportantError struct {
	err   error
	cause error
}

func NewVeryImportantError(cause error) VeryImportantError {
	return VeryImportantError{err: errors.New("oчень важная ошибка"), cause: cause}
}

func (err VeryImportantError) Error() string {
	return err.err.Error()
}

func (err VeryImportantError) Unwrap() error {
	return err.cause
}
