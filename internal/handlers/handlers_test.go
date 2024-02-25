package handlers

import (
	"os"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func TestMain(m *testing.M) {
	e = echo.New()
	e.Validator = &Validator{validator: validator.New()}
	code := m.Run()
	os.Exit(code)
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
