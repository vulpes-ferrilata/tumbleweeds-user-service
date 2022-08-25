package documents

type User struct {
	DocumentRoot `bson:",inline"`
	DisplayName  string `bson:"display_name"`
}
