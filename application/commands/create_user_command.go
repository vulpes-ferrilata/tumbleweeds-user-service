package commands

type CreateUserCommand struct {
	ID          string `validate:"required,uuid4"`
	DisplayName string `validate:"required,lte=20"`
}
