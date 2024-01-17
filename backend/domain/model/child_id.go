package model

import "github.com/google/uuid"

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GO_PACKAGE

type ChildID interface {
	String() string
}

type childID struct {
	id string
}

func NewChildID() (ChildID, error) {
	id := uuid.New()
	return &childID{id: id.String()}, nil
}

func NewChildIDFromString(id string) ChildID {
	return &childID{
		id: id,
	}
}

func (c *childID) String() string {
	return c.id
}
