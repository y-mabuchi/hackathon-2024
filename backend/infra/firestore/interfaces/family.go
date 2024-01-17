package interfaces

import "github.com/y-mabuchi/child-memories/backend/domain/model"

type FamilyType struct {
	ID   string `firestore:"id"`
	Name string `firestore:"name"`
}

func (f FamilyType) Model() model.Family {
	return model.NewFamilyFromRepository(f.ID, f.Name)
}

func NewFamilyFromModel(family model.Family) *FamilyType {
	return &FamilyType{
		ID:   family.ID().String(),
		Name: family.Name().String(),
	}
}
