package model

import "time"

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GO_PACKAGE

type ChildProfile interface {
	ID() ChildID
	FamilyID() FamilyID
	Name() ChildName
	Gender() ChildGender
	Birthday() ChildBirthday
	ChangeFirstName(firstname string) error
	ChangeLastName(lastname string) error
	ChangeGender(gender ChildGender) error
	ChangeBirthday(birthday time.Time) error
}

type childProfile struct {
	id       ChildID
	familyID FamilyID
	name     ChildName
	gender   ChildGender
	birthday ChildBirthday
}

func NewChildProfile(
	familyID FamilyID,
	name ChildName,
	gender ChildGender,
	birthday ChildBirthday,
) (ChildProfile, error) {
	id, err := NewChildID()
	if err != nil {
		return nil, err
	}

	return &childProfile{
		id:       id,
		familyID: familyID,
		name:     name,
		gender:   gender,
		birthday: birthday,
	}, nil
}

func NewChildProfileFromRepository(
	id, familyID, firstname, lastname, gender string,
	birthday time.Time,
) ChildProfile {
	cID := NewChildIDFromString(id)
	fID := NewFamilyIDFromString(familyID)
	cName, _ := NewChildName(firstname, lastname)
	cBirthday, _ := NewChildBirthday(birthday)
	cGender, _ := ParseChildGender(gender)

	return &childProfile{
		id:       cID,
		familyID: fID,
		name:     cName,
		gender:   cGender,
		birthday: cBirthday,
	}
}

func (c *childProfile) ID() ChildID {
	return c.id
}

func (c *childProfile) FamilyID() FamilyID {
	return c.familyID
}

func (c *childProfile) Name() ChildName {
	return c.name
}

func (c *childProfile) Gender() ChildGender {
	return c.gender
}

func (c *childProfile) Birthday() ChildBirthday {
	return c.birthday
}

func (c *childProfile) ChangeFirstName(firstname string) error {
	cName, err := NewChildName(firstname, c.name.LastName())
	if err != nil {
		return err
	}

	c.name = cName

	return nil
}

func (c *childProfile) ChangeLastName(lastname string) error {
	cName, err := NewChildName(c.name.FirstName(), lastname)
	if err != nil {
		return err
	}

	c.name = cName

	return nil
}

func (c *childProfile) ChangeGender(gender ChildGender) error {
	c.gender = gender

	return nil
}

func (c *childProfile) ChangeBirthday(birthday time.Time) error {
	cBirthday, err := NewChildBirthday(birthday)
	if err != nil {
		return err
	}

	c.birthday = cBirthday

	return nil
}
