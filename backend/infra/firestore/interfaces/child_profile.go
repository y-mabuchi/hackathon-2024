package interfaces

import (
	"time"

	"github.com/y-mabuchi/child-memories/backend/domain/model"
)

type ChildProfileType struct {
	ID        string    `firestore:"id"`
	FirstName string    `firestore:"first_name"`
	LastName  string    `firestore:"last_name"`
	Gender    string    `firestore:"gender"`
	Birthday  time.Time `firestore:"birthday"`
	FamilyID  string    `firestore:"family_id"`
}

func (c *ChildProfileType) ConvertToModel() model.ChildProfile {
	return model.NewChildProfileFromRepository(
		c.ID,
		c.FamilyID,
		c.FirstName,
		c.LastName,
		c.Gender,
		c.Birthday,
	)
}

func NewChildProfileFromModel(profile model.ChildProfile) *ChildProfileType {
	return &ChildProfileType{
		ID:        profile.ID().String(),
		FamilyID:  profile.FamilyID().String(),
		FirstName: profile.Name().FirstName(),
		LastName:  profile.Name().LastName(),
		Gender:    string(profile.Gender()),
		Birthday:  profile.Birthday().ToTime(),
	}
}
