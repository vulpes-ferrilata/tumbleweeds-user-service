package commands

type CreateUserCommand struct {
	ID          string `validate:"required,objectid"`
	DisplayName string `validate:"required,gte=1,lte=20"`
}
