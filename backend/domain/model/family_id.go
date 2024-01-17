package model

import "github.com/google/uuid"

type FamilyID interface {
	String() string
}

type familyID struct {
	id string
}

func NewFamilyID() (FamilyID, error) {
	id := uuid.New()
	return &familyID{
		id: id.String(),
	}, nil
}

func (f *familyID) String() string {
	return f.id
}

func NewFamilyIDFromString(id string) FamilyID {
	return &familyID{
		id: id,
	}
}
