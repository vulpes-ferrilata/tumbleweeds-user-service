package commands

type CreateUserCommand struct {
	ID          string `validate:"required,objectid"`
	DisplayName string `validate:"required,min=1,max=20"`
}
