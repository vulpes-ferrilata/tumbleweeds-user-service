package commands

type CreateUser struct {
	UserID      string `validate:"required,objectid"`
	DisplayName string `validate:"required,min=1,max=20"`
}
