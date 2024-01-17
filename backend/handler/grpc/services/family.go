package services

import (
	"context"
	"log"

	"github.com/y-mabuchi/child-memories/backend/handler/grpc/gerror"

	"github.com/y-mabuchi/child-memories/backend/usecase"

	"github.com/y-mabuchi/child-memories/backend/usecase/input"

	"github.com/y-mabuchi/child-memories/backend/handler/grpc/generated"
	"github.com/y-mabuchi/child-memories/backend/registry"
	"github.com/y-mabuchi/child-memories/backend/usecase/dto"
)

type FamilyService struct {
	familyUseCase usecase.FamilyUseCase
	generated.UnimplementedFamilyServiceServer
}

func NewFamilyService(regi registry.Registry) *FamilyService {
	return &FamilyService{
		familyUseCase: regi.NewFamilyUseCase(),
	}
}

func NewFamilyFromDTO(family *dto.FamilyDTO) *generated.Family {
	return &generated.Family{
		Id:   family.ID,
		Name: family.Name,
	}
}

func (s *FamilyService) FindOneFamily(
	ctx context.Context,
	req *generated.FindOneFamilyRequest,
) (*generated.FindOneFamilyResponse, error) {
	familyDTO, err := s.familyUseCase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.NewErr(err)
	}

	gFamily := NewFamilyFromDTO(familyDTO)

	return &generated.FindOneFamilyResponse{
		Family: gFamily,
	}, nil
}

func (s *FamilyService) CreateFamily(
	ctx context.Context,
	req *generated.CreateFamilyRequest,
) (*generated.CreateFamilyResponse, error) {
	log.Println("create family start.")
	createInput := input.FamilyCreateInput{
		Name: req.Name,
	}

	familyDTO, err := s.familyUseCase.Create(ctx, createInput)
	if err != nil {
		return nil, gerror.NewErr(err)
	}

	gProfile := NewFamilyFromDTO(familyDTO)

	return &generated.CreateFamilyResponse{
		Family: gProfile,
	}, nil
}

func (s *FamilyService) UpdateFamily(
	ctx context.Context,
	req *generated.UpdateFamilyRequest,
) (*generated.UpdateFamilyResponse, error) {
	updateInput := input.FamilyUpdateInput{
		ID:   req.Family.Id,
		Name: req.Family.Name,
	}

	familyDTO, err := s.familyUseCase.Update(ctx, updateInput)
	if err != nil {
		return nil, gerror.NewErr(err)
	}

	gFamily := NewFamilyFromDTO(familyDTO)

	return &generated.UpdateFamilyResponse{
		Family: gFamily,
	}, nil
}

func (s *FamilyService) DeleteFamily(
	ctx context.Context,
	req *generated.DeleteFamilyRequest,
) (*generated.DeleteFamilyResponse, error) {
	deleteInput := input.FamilyDeleteInput{
		ID: req.Id,
	}

	if err := s.familyUseCase.Delete(ctx, deleteInput); err != nil {
		return nil, gerror.NewErr(err)
	}

	return &generated.DeleteFamilyResponse{
		Success: true,
	}, nil
}
