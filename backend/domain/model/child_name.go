package model

import "fmt"

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GO_PACKAGE

type ChildName interface {
	FullName() string
	FirstName() string
	LastName() string
}

type childName struct {
	firstName string
	lastName  string
}

func NewChildName(firstName, lastName string) (ChildName, error) {
	return &childName{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

func (c *childName) FullName() string {
	return fmt.Sprintf("%s %s", c.firstName, c.lastName)
}

func (c *childName) FirstName() string {
	return c.firstName
}

func (c *childName) LastName() string {
	return c.lastName
}
