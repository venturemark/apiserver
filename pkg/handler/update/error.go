package update

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var invalidConfigError = &tracer.Error{
	Kind: "invalidConfigError",
}

func IsInvalidConfig(err error) bool {
	return errors.Is(err, invalidConfigError)
}

var invalidInputError = &tracer.Error{
	Kind: "invalidInputError",
	Desc: "Could not identify the desired operation based on the provided input.",
}

func IsInvalidInput(err error) bool {
	return errors.Is(err, invalidInputError)
}
