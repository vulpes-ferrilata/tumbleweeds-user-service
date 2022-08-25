package commands

type CreateUserCommand struct {
	ID          string `validate:"required,objectid"`
	DisplayName string `validate:"required,lte=20"`
}
