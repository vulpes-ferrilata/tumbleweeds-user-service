package app_errors

var (
	ErrUserNotFound           error = NewNotFoundError("user")
	ErrUserStateInconsistence error = NewInconsistenceStateError("user")
	ErrStaleUser              error = NewStaleObjectError("user")
)
