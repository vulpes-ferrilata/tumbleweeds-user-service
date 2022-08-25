package app_errors

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewInconsistenceStateError(object string) AppError {
	return &inconsistenceStateError{
		object: object,
	}
}

type inconsistenceStateError struct {
	object string
}

func (s inconsistenceStateError) Error() string {
	return fmt.Sprintf("%s state does not match", s.object)
}

func (s inconsistenceStateError) Status(translator ut.Translator) *status.Status {
	object, err := translator.T(s.object)
	if err != nil {
		object = s.object
	}

	detail, err := translator.T("inconsistence-state-error", object)
	if err != nil {
		detail = s.Error()
	}

	stt := status.New(codes.FailedPrecondition, detail)

	return stt
}
