package repository

import (
	"context"

	"github.com/y-mabuchi/child-memories/backend/domain/model"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GO_PACKAGE

type ChildProfileRepository interface {
	FindByID(ctx context.Context, cID model.ChildID) (model.ChildProfile, error)
	Create(ctx context.Context, profile model.ChildProfile) (model.ChildProfile, error)
	Update(ctx context.Context, profile model.ChildProfile) (model.ChildProfile, error)
	Delete(ctx context.Context, cID model.ChildID) error
}
