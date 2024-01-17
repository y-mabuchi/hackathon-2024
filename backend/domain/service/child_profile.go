package service

import (
	"context"

	"github.com/y-mabuchi/child-memories/backend/domain/model"
	"github.com/y-mabuchi/child-memories/backend/domain/repository"
)

type ChildProfileService interface {
	Exists(ctx context.Context, id model.ChildID) bool
}

type childProfileService struct {
	childProfileRepository repository.ChildProfileRepository
}

func NewChildProfileService(childProfileRepository repository.ChildProfileRepository) ChildProfileService {
	return &childProfileService{
		childProfileRepository: childProfileRepository,
	}
}

func (c *childProfileService) Exists(ctx context.Context, id model.ChildID) bool {
	_, err := c.childProfileRepository.FindByID(ctx, id)
	if err != nil {
		return false
	}

	return true
}
