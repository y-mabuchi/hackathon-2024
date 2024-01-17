package service

import (
	"context"
	"log"

	"github.com/y-mabuchi/child-memories/backend/domain/model"
	"github.com/y-mabuchi/child-memories/backend/domain/repository"
)

type FamilyService interface {
	Exists(ctx context.Context, id model.FamilyID) bool
}

type familyService struct {
	familyRepository repository.FamilyRepository
}

func NewFamilyService(familyRepository repository.FamilyRepository) FamilyService {
	return &familyService{
		familyRepository: familyRepository,
	}
}

func (f *familyService) Exists(ctx context.Context, id model.FamilyID) bool {
	_, err := f.familyRepository.FindByID(ctx, id)
	if err != nil {
		log.Printf("family find by id err: %v", err)
		return false
	}

	return true
}
