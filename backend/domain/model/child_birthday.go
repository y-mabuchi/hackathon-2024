package model

import "time"

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GO_PACKAGE

type ChildBirthday interface {
	ToTime() time.Time
}

type childBirthday struct {
	birthday time.Time
}

func NewChildBirthday(birthday time.Time) (ChildBirthday, error) {
	return &childBirthday{birthday: birthday}, nil
}

func (c *childBirthday) ToTime() time.Time {
	return c.birthday
}
