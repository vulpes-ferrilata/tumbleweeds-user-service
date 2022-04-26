package context

import "github.com/pkg/errors"

var (
	ErrContextValueNotFound error = errors.New("context value not found")
)
