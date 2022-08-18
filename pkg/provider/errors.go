package provider

import (
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

var (
	// ErrInternal reports an internal error.
	ErrInternal = errors.New("internal error")
	// ErrNotFound reports if entity is not found.
	ErrNotFound = errors.New("not found")
)
