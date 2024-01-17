package usecase

import (
	"context"

	"github.com/y-mabuchi/child-memories/backend/domain/model"
	"github.com/y-mabuchi/child-memories/backend/domain/repository"
	"github.com/y-mabuchi/child-memories/backend/domain/service"
	"github.com/y-mabuchi/child-memories/backend/usecase/dto"
	"github.com/y-mabuchi/child-memories/backend/usecase/input"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FamilyUseCase interface {
	GetByID(ctx context.Context, id string) (*dto.FamilyDTO, error)
	Create(ctx context.Context, input input.FamilyCreateInput) (*dto.FamilyDTO, error)
	Update(ctx context.Context, input input.FamilyUpdateInput) (*dto.FamilyDTO, error)
	Delete(ctx context.Context, input input.FamilyDeleteInput) error
}

type familyUseCase struct {
	familyRepository repository.FamilyRepository
	familyService    service.FamilyService
}

func NewFamilyUseCase(
	familyRepository repository.FamilyRepository,
	familyService service.FamilyService,
) FamilyUseCase {
	return &familyUseCase{
		familyRepository: familyRepository,
		familyService:    familyService,
	}
}

func (f *familyUseCase) GetByID(
	ctx context.Context,
	id string,
) (*dto.FamilyDTO, error) {
	fID := model.NewFamilyIDFromString(id)

	rFamily, err := f.familyRepository.FindByID(ctx, fID)
	if status.Code(err) == codes.NotFound {
		return nil, model.NewErrNotFound("family", fID.String())
	}
	if err != nil {
		return nil, err
	}

	return dto.NewFamilyDTOFromModel(rFamily), nil
}

func (f *familyUseCase) Create(
	ctx context.Context,
	input input.FamilyCreateInput,
) (*dto.FamilyDTO, error) {
	fName, err := model.NewFamilyName(input.Name)
	if err != nil {
		return nil, err
	}

	mFamily, err := model.NewFamily(fName)
	if err != nil {
		return nil, err
	}

	rFamily, err := f.familyRepository.Create(ctx, mFamily)
	if err != nil {
		return nil, err
	}

	return dto.NewFamilyDTOFromModel(rFamily), nil
}

func (f *familyUseCase) Update(
	ctx context.Context,
	input input.FamilyUpdateInput,
) (*dto.FamilyDTO, error) {
	fID := model.NewFamilyIDFromString(input.ID)
	if ok := f.familyService.Exists(ctx, fID); !ok {
		return nil, model.NewErrNotFound("family", fID.String())
	}

	mFamily, err := f.familyRepository.FindByID(ctx, fID)
	if err != nil {
		return nil, err
	}

	if err = mFamily.ChangeName(input.Name); err != nil {
		return nil, err
	}

	family, err := f.familyRepository.Update(ctx, mFamily)
	if err != nil {
		return nil, err
	}

	return dto.NewFamilyDTOFromModel(family), nil
}

func (f *familyUseCase) Delete(
	ctx context.Context,
	input input.FamilyDeleteInput,
) error {
	fID := model.NewFamilyIDFromString(input.ID)
	if ok := f.familyService.Exists(ctx, fID); !ok {
		return model.NewErrNotFound("family", fID.String())
	}

	if err := f.familyRepository.Delete(ctx, fID); err != nil {
		return err
	}

	return nil
}
