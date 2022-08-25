package queries

type GetUserByIDQuery struct {
	ID string `validate:"required,objectid"`
}
