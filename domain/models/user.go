package models

type User struct {
	aggregateRoot
	displayName string
}

func (u User) GetDisplayName() string {
	return u.displayName
}
