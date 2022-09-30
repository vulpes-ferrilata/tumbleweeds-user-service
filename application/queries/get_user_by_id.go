package queries

type GetUserByID struct {
	UserID string `validate:"required,objectid"`
}
