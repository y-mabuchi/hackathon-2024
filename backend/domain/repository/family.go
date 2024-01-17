package repository

import (
	"context"

	"github.com/y-mabuchi/child-memories/backend/domain/model"
)

type FamilyRepository interface {
	FindByID(ctx context.Context, id model.FamilyID) (model.Family, error)
	Create(ctx context.Context, family model.Family) (model.Family, error)
	Update(ctx context.Context, family model.Family) (model.Family, error)
	Delete(ctx context.Context, id model.FamilyID) error
}
