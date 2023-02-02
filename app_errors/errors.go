package app_errors

var (
	ErrUserNotFound error = NewNotFoundError("user")
	ErrStaleUser    error = NewStaleObjectError("user")
)
