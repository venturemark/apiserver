package validate

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var invalidInputError = &tracer.Error{
	Kind: "invalidInputError",
}

func IsInvalidInput(err error) bool {
	return errors.Is(err, invalidInputError)
}
