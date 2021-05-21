package common

import "github.com/pkg/errors"

// These error codes are common errors happened in most packages.
var (
	ErrNotImpl = errors.Errorf("not implemented")
)
