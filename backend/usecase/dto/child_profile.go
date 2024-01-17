package dto

import (
	"time"

	"github.com/y-mabuchi/child-memories/backend/domain/model"
)

type ChildProfileDTO struct {
	ID        string
	FamilyID  string
	FirstName string
	LastName  string
	Gender    string
	Birthday  time.Time
}

func NewChildProfileDTO(cProfile model.ChildProfile) *ChildProfileDTO {
	return &ChildProfileDTO{
		ID:        cProfile.ID().String(),
		FamilyID:  cProfile.FamilyID().String(),
		FirstName: cProfile.Name().FirstName(),
		LastName:  cProfile.Name().LastName(),
		Gender:    string(cProfile.Gender()),
		Birthday:  cProfile.Birthday().ToTime(),
	}
}
