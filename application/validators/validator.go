package validators

import "context"

type Validator[T any] interface {
	Validate(ctx context.Context, input T) error
}
