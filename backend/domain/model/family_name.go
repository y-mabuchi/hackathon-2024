package model

type FamilyName interface {
	String() string
}

type familyName struct {
	name string
}

func NewFamilyName(name string) (FamilyName, error) {
	return &familyName{
		name: name,
	}, nil
}

func (f *familyName) String() string {
	return f.name
}
