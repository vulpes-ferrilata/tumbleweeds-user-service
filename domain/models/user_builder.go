package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserBuilder interface {
	SetID(id primitive.ObjectID) UserBuilder
	SetDisplayName(displayName string) UserBuilder
	SetVersion(version int) UserBuilder
	Create() *User
}

func NewUserBuilder() UserBuilder {
	return &userBuilder{}
}

type userBuilder struct {
	id          primitive.ObjectID
	displayName string
	version     int
}

func (u *userBuilder) SetID(id primitive.ObjectID) UserBuilder {
	u.id = id

	return u
}

func (u *userBuilder) SetDisplayName(displayName string) UserBuilder {
	u.displayName = displayName

	return u
}

func (u *userBuilder) SetVersion(version int) UserBuilder {
	u.version = version

	return u
}

func (u userBuilder) Create() *User {
	return &User{
		aggregateRoot: aggregateRoot{
			aggregate: aggregate{
				id: u.id,
			},
			version: u.version,
		},
		displayName: u.displayName,
	}
}
