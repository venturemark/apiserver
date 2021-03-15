package user

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
	Desc: "This error indicates a problem with the request payload. The desired operation could not be identified based on the provided user input.",
}

func IsInvalidInput(err error) bool {
	return errors.Is(err, invalidInputError)
}

var invalidUserError = &tracer.Error{
	Kind: "invalidUserError",
	Desc: "This error indicates a problem with the request authentication. A jwt token must be present in the grpc metadata. The access token must be provided using the bearer scheme",
}

func IsInvalidUser(err error) bool {
	return errors.Is(err, invalidUserError)
}
