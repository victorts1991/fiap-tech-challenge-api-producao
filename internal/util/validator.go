package util

import "github.com/go-playground/validator/v10"

//go:generate mockgen -source=$GOFILE -package=mock_util -destination=../../test/mock/util/$GOFILE

type Validator interface {
	ValidateStruct(i interface{}) error
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() Validator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) ValidateStruct(i interface{}) error {
	return cv.validator.Struct(i)
}
