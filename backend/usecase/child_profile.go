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

type ChildProfileUseCase interface {
	GetByID(ctx context.Context, id, family_id string) (*dto.ChildProfileDTO, error)
	Create(ctx context.Context, input input.ChildProfileCreateInput) (*dto.ChildProfileDTO, error)
	Update(ctx context.Context, input input.ChildProfileUpdateInput) (*dto.ChildProfileDTO, error)
	Delete(ctx context.Context, input input.ChildProfileDeleteInput) error
}

type childProfileUseCase struct {
	childProfileRepo    repository.ChildProfileRepository
	childProfileService service.ChildProfileService
	familyService       service.FamilyService
}

func NewChildProfileUseCase(
	childProfileRepository repository.ChildProfileRepository,
	childProfileService service.ChildProfileService,
	familyService service.FamilyService,
) ChildProfileUseCase {
	return &childProfileUseCase{
		childProfileRepo:    childProfileRepository,
		childProfileService: childProfileService,
		familyService:       familyService,
	}
}

func (c *childProfileUseCase) GetByID(
	ctx context.Context,
	id, familyID string,
) (*dto.ChildProfileDTO, error) {
	fID := model.NewFamilyIDFromString(familyID)
	if ok := c.familyService.Exists(ctx, fID); !ok {
		return nil, model.NewErrNotFound("family", fID.String())
	}

	cID := model.NewChildIDFromString(id)

	cProfile, err := c.childProfileRepo.FindByID(ctx, cID)
	if status.Code(err) == codes.NotFound {
		return nil, model.NewErrNotFound("child profile", cID.String())
	}
	if err != nil {
		return nil, err
	}

	return dto.NewChildProfileDTO(cProfile), nil
}

func (c *childProfileUseCase) Create(
	ctx context.Context,
	input input.ChildProfileCreateInput,
) (*dto.ChildProfileDTO, error) {
	fID := model.NewFamilyIDFromString(input.FamilyID)
	if ok := c.familyService.Exists(ctx, fID); !ok {
		return nil, model.NewErrNotFound("family", fID.String())
	}

	cName, err := model.NewChildName(input.FirstName, input.LastName)
	if err != nil {
		return nil, err
	}

	cGender, err := model.ParseChildGender(input.Gender)
	if err != nil {
		return nil, err
	}

	cBirthday, err := model.NewChildBirthday(input.Birthday)
	if err != nil {
		return nil, err
	}

	cProfile, err := model.NewChildProfile(fID, cName, cGender, cBirthday)
	if err != nil {
		return nil, err
	}

	cProfile, err = c.childProfileRepo.Create(ctx, cProfile)
	if err != nil {
		return nil, err
	}

	return dto.NewChildProfileDTO(cProfile), nil
}

func (c *childProfileUseCase) Update(
	ctx context.Context,
	input input.ChildProfileUpdateInput,
) (*dto.ChildProfileDTO, error) {
	fID := model.NewFamilyIDFromString(input.FamilyID)
	if ok := c.familyService.Exists(ctx, fID); !ok {
		return nil, model.NewErrNotFound("family", fID.String())
	}

	cID := model.NewChildIDFromString(input.ID)
	if ok := c.childProfileService.Exists(ctx, cID); !ok {
		return nil, model.NewErrNotFound("child profile", cID.String())
	}

	cProfile, err := c.childProfileRepo.FindByID(ctx, cID)
	if err != nil {
		return nil, err
	}

	if err = cProfile.ChangeFirstName(input.FirstName); err != nil {
		return nil, err
	}

	if err = cProfile.ChangeLastName(input.LastName); err != nil {
		return nil, err
	}

	cGender, err := model.ParseChildGender(input.Gender)
	if err != nil {
		return nil, err
	}
	if err = cProfile.ChangeGender(cGender); err != nil {
		return nil, err
	}

	if err = cProfile.ChangeBirthday(input.Birthday); err != nil {
		return nil, err
	}

	cProfile, err = c.childProfileRepo.Update(ctx, cProfile)
	if err != nil {
		return nil, err
	}

	return dto.NewChildProfileDTO(cProfile), nil
}

func (c *childProfileUseCase) Delete(ctx context.Context, input input.ChildProfileDeleteInput) error {
	fID := model.NewFamilyIDFromString(input.FamilyID)
	if ok := c.familyService.Exists(ctx, fID); !ok {
		return model.NewErrNotFound("family", fID.String())
	}

	cID := model.NewChildIDFromString(input.ID)
	if ok := c.childProfileService.Exists(ctx, cID); !ok {
		return model.NewErrNotFound("child profile", cID.String())
	}

	if err := c.childProfileRepo.Delete(ctx, cID); err != nil {
		return err
	}

	return nil
}
