package dto

import "github.com/y-mabuchi/child-memories/backend/domain/model"

type FamilyDTO struct {
	ID   string
	Name string
}

func NewFamilyDTOFromModel(family model.Family) *FamilyDTO {
	return &FamilyDTO{
		ID:   family.ID().String(),
		Name: family.Name().String(),
	}
}
