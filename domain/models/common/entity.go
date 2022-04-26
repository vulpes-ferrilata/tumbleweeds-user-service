package common

import "github.com/google/uuid"

func NewEntity(id uuid.UUID, version int) Entity {
	entity := Entity{
		id:      id,
		version: version,
	}
	return entity
}

type Entity struct {
	id      uuid.UUID
	version int
}

func (e Entity) GetID() uuid.UUID {
	return e.id
}

func (e Entity) GetVersion() int {
	return e.version
}
