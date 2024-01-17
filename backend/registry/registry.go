package registry

import (
	"cloud.google.com/go/firestore"
	"github.com/y-mabuchi/child-memories/backend/domain/repository"
	"github.com/y-mabuchi/child-memories/backend/domain/service"
	irepository "github.com/y-mabuchi/child-memories/backend/infra/firestore/repository"
	"github.com/y-mabuchi/child-memories/backend/usecase"
)

type Registry interface {
	NewChildProfileRepository() repository.ChildProfileRepository
	NewChildProfileUseCase() usecase.ChildProfileUseCase
	NewFamilyRepository() repository.FamilyRepository
	NewFamilyUseCase() usecase.FamilyUseCase
}

type registry struct {
	client *firestore.Client
}

func NewRegistry(client *firestore.Client) Registry {
	return &registry{
		client: client,
	}
}

func (r *registry) NewChildProfileRepository() repository.ChildProfileRepository {
	return irepository.NewChildProfileRepository(r.client)
}

func (r *registry) NewChildProfileUseCase() usecase.ChildProfileUseCase {
	cRepo := r.NewChildProfileRepository()
	cSvc := service.NewChildProfileService(cRepo)
	fRepo := r.NewFamilyRepository()
	fSvc := service.NewFamilyService(fRepo)

	return usecase.NewChildProfileUseCase(
		cRepo,
		cSvc,
		fSvc,
	)
}

func (r *registry) NewFamilyRepository() repository.FamilyRepository {
	return irepository.NewFamilyRepository(r.client)
}

func (r *registry) NewFamilyUseCase() usecase.FamilyUseCase {
	fRepo := r.NewFamilyRepository()
	fSvc := service.NewFamilyService(fRepo)

	return usecase.NewFamilyUseCase(
		fRepo,
		fSvc,
	)
}
