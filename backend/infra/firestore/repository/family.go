package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/y-mabuchi/child-memories/backend/domain/model"
	"github.com/y-mabuchi/child-memories/backend/domain/repository"
	"github.com/y-mabuchi/child-memories/backend/infra/firestore/interfaces"
)

const (
	FamilyPath = "families"
)

type familyRepository struct {
	client *firestore.Client
}

func NewFamilyRepository(
	client *firestore.Client,
) repository.FamilyRepository {
	return &familyRepository{
		client: client,
	}
}

func (f *familyRepository) FindByID(
	ctx context.Context,
	id model.FamilyID,
) (model.Family, error) {
	docSnap, err := f.client.
		Collection(FamilyPath).
		Doc(id.String()).
		Get(ctx)
	if err != nil {
		return nil, err
	}

	var family interfaces.FamilyType
	if err = docSnap.DataTo(&family); err != nil {
		return nil, err
	}

	return family.Model(), nil
}

func (f *familyRepository) Create(
	ctx context.Context,
	family model.Family,
) (model.Family, error) {
	_, err := f.client.
		Collection(FamilyPath).
		Doc(family.ID().String()).
		Set(
			ctx,
			interfaces.NewFamilyFromModel(family),
		)
	if err != nil {
		return nil, err
	}

	rFamily, err := f.FindByID(ctx, family.ID())
	if err != nil {
		return nil, err
	}

	return rFamily, nil
}

func (f *familyRepository) Update(
	ctx context.Context,
	family model.Family,
) (model.Family, error) {
	_, err := f.client.
		Collection(FamilyPath).
		Doc(family.ID().String()).
		Set(ctx,
			interfaces.NewFamilyFromModel(family),
		)
	if err != nil {
		return nil, err
	}

	rFamily, err := f.FindByID(ctx, family.ID())
	if err != nil {
		return nil, err
	}

	return rFamily, nil
}

func (f *familyRepository) Delete(
	ctx context.Context,
	id model.FamilyID,
) error {
	_, err := f.client.
		Collection(FamilyPath).
		Doc(id.String()).
		Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
