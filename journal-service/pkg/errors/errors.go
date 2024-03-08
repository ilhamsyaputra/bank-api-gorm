package Errors

import "errors"

var (
	RequestValidation = NewError("gagal validasi request")
)

func NewError(message string) error {
	return errors.New(message)
}
