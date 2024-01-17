package model

type Family interface {
	ID() FamilyID
	Name() FamilyName
	ChangeName(name string) error
}

type family struct {
	id   FamilyID
	name FamilyName
}

func NewFamily(name FamilyName) (Family, error) {
	id, err := NewFamilyID()
	if err != nil {
		return nil, err
	}

	return &family{
		id:   id,
		name: name,
	}, nil
}

func NewFamilyFromRepository(id, name string) Family {
	fID := NewFamilyIDFromString(id)
	fName, _ := NewFamilyName(name)

	return &family{
		id:   fID,
		name: fName,
	}
}

func (f *family) ID() FamilyID {
	return f.id
}

func (f *family) Name() FamilyName {
	return f.name
}

func (f *family) ChangeName(name string) error {
	fName, err := NewFamilyName(name)
	if err != nil {
		return err
	}

	f.name = fName

	return nil
}
