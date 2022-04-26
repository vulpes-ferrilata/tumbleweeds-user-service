package service

func NewBusinessRuleError(message string) error {
	return &BusinessRuleError{
		msg: message,
	}
}

type BusinessRuleError struct {
	msg string
}

func (b BusinessRuleError) Error() string {
	return b.msg
}
