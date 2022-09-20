package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserBuilder struct {
	id          primitive.ObjectID
	displayName string
	version     int
}

func (u UserBuilder) SetID(id primitive.ObjectID) UserBuilder {
	u.id = id

	return u
}

func (u UserBuilder) SetDisplayName(displayName string) UserBuilder {
	u.displayName = displayName

	return u
}

func (u UserBuilder) SetVersion(version int) UserBuilder {
	u.version = version

	return u
}

func (u UserBuilder) Create() *User {
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
