package services

import (
	"context"

	"github.com/y-mabuchi/child-memories/backend/handler/grpc/gerror"

	"github.com/y-mabuchi/child-memories/backend/usecase"

	"github.com/y-mabuchi/child-memories/backend/handler/grpc/generated"
	"github.com/y-mabuchi/child-memories/backend/registry"
	"github.com/y-mabuchi/child-memories/backend/usecase/dto"
	"github.com/y-mabuchi/child-memories/backend/usecase/input"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChildProfileService struct {
	childProfileUseCase usecase.ChildProfileUseCase
	generated.UnimplementedChildProfileServiceServer
}

func NewChildProfileService(regi registry.Registry) *ChildProfileService {
	return &ChildProfileService{
		childProfileUseCase: regi.NewChildProfileUseCase(),
	}
}

func GenderFromString(gender string) generated.ChildGender {
	switch gender {
	case "CHILD_GENDER_BOY":
		return generated.ChildGender_CHILD_GENDER_BOY
	case "CHILD_GENDER_GIRL":
		return generated.ChildGender_CHILD_GENDER_GIRL
	default:
		return generated.ChildGender_CHILD_GENDER_UNSPECIFIED
	}
}

func NewChildProfileFromDTO(cProfile *dto.ChildProfileDTO) *generated.ChildProfile {
	return &generated.ChildProfile{
		Id:        cProfile.ID,
		FirstName: cProfile.FirstName,
		LastName:  cProfile.LastName,
		Gender:    GenderFromString(cProfile.Gender),
		Birthday:  timestamppb.New(cProfile.Birthday),
		FamilyId:  cProfile.FamilyID,
	}
}

func (s *ChildProfileService) FindOneChildProfile(
	ctx context.Context,
	req *generated.FindOneChildProfileRequest,
) (*generated.FindOneChildProfileResponse, error) {
	cProfile, err := s.childProfileUseCase.GetByID(ctx, req.Id, req.FamilyId)
	if err != nil {
		return nil, gerror.NewErr(err)
	}

	gProfile := NewChildProfileFromDTO(cProfile)

	return &generated.FindOneChildProfileResponse{
		Profile: gProfile,
	}, nil
}

func (s *ChildProfileService) CreateChildProfile(
	ctx context.Context,
	req *generated.CreateChildProfileRequest,
) (*generated.CreateChildProfileResponse, error) {
	createInput := input.ChildProfileCreateInput{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    req.Gender.String(),
		Birthday:  req.Birthday.AsTime(),
		FamilyID:  req.FamilyId,
	}

	cProfile, err := s.childProfileUseCase.Create(ctx, createInput)
	if err != nil {
		return nil, gerror.NewErr(err)
	}

	gProfile := NewChildProfileFromDTO(cProfile)

	return &generated.CreateChildProfileResponse{
		Profile: gProfile,
	}, nil
}

func (s *ChildProfileService) UpdateChildProfile(
	ctx context.Context,
	req *generated.UpdateChildProfileRequest,
) (*generated.UpdateChildProfileResponse, error) {
	updateInput := input.ChildProfileUpdateInput{
		ID:        req.Profile.Id,
		FirstName: req.Profile.FirstName,
		LastName:  req.Profile.LastName,
		Gender:    req.Profile.Gender.String(),
		Birthday:  req.Profile.Birthday.AsTime(),
		FamilyID:  req.Profile.FamilyId,
	}

	cProfile, err := s.childProfileUseCase.Update(ctx, updateInput)
	if err != nil {
		return nil, gerror.NewErr(err)
	}

	gProfile := NewChildProfileFromDTO(cProfile)

	return &generated.UpdateChildProfileResponse{
		Profile: gProfile,
	}, nil
}

func (s *ChildProfileService) DeleteChildProfile(
	ctx context.Context,
	req *generated.DeleteChildProfileRequest,
) (*generated.DeleteChildProfileResponse, error) {
	deleteInput := input.ChildProfileDeleteInput{
		ID:       req.Id,
		FamilyID: req.FamilyId,
	}

	if err := s.childProfileUseCase.Delete(ctx, deleteInput); err != nil {
		return &generated.DeleteChildProfileResponse{
			Success: false,
		}, gerror.NewErr(err)
	}

	return &generated.DeleteChildProfileResponse{
		Success: true,
	}, nil
}
