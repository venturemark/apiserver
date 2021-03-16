package association

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

var invalidMetadataError = &tracer.Error{
	Kind: "invalidMetadataError",
}

func IsInvalidMetadata(err error) bool {
	return errors.Is(err, invalidMetadataError)
}
