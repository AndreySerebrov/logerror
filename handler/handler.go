package handler

import (
	"err_test/apperrors"
	"errors"
	"fmt"
)

type Doer interface {
	Do() error
}

type Handler struct {
	doer Doer
}

func NewHandler(doer Doer) Handler {
	return Handler{doer: doer}
}

func (h Handler) Handle() error {
	err := h.doer.Do()
	if errors.As(err, &apperrors.ImportantError{}) {
		return fmt.Errorf("важная ошибка")
	}
	if errors.As(err, &apperrors.VeryImportantError{}) {
		return fmt.Errorf("очень важная ошибка")
	}
	return nil
}
