package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/y-mabuchi/child-memories/backend/domain/model"
	"github.com/y-mabuchi/child-memories/backend/domain/repository"
	"github.com/y-mabuchi/child-memories/backend/infra/firestore/interfaces"
)

const (
	ChildProfilePath = "child_profiles"
)

type childProfileRepository struct {
	client *firestore.Client
}

func NewChildProfileRepository(client *firestore.Client) repository.ChildProfileRepository {
	return &childProfileRepository{
		client: client,
	}
}

func (c *childProfileRepository) FindByID(
	ctx context.Context,
	cID model.ChildID,
) (model.ChildProfile, error) {
	docSnap, err := c.client.
		Collection(ChildProfilePath).
		Doc(cID.String()).
		Get(ctx)
	if err != nil {
		return nil, err
	}

	var profile interfaces.ChildProfileType
	if err = docSnap.DataTo(&profile); err != nil {
		return nil, err
	}

	return profile.ConvertToModel(), nil
}

func (c *childProfileRepository) Create(
	ctx context.Context,
	profile model.ChildProfile,
) (model.ChildProfile, error) {
	_, err := c.client.
		Collection(ChildProfilePath).
		Doc(profile.ID().String()).
		Set(
			ctx,
			interfaces.NewChildProfileFromModel(profile),
		)
	if err != nil {
		return nil, err
	}

	cProfile, err := c.FindByID(ctx, profile.ID())
	if err != nil {
		return nil, err
	}

	return cProfile, nil
}

func (c *childProfileRepository) Update(
	ctx context.Context,
	profile model.ChildProfile,
) (model.ChildProfile, error) {
	_, err := c.client.
		Collection(ChildProfilePath).
		Doc(profile.ID().String()).
		Set(
			ctx,
			interfaces.NewChildProfileFromModel(profile),
		)
	if err != nil {
		return nil, err
	}

	cProfile, err := c.FindByID(ctx, profile.ID())
	if err != nil {
		return nil, err
	}

	return cProfile, nil
}

func (c *childProfileRepository) Delete(
	ctx context.Context,
	id model.ChildID,
) error {
	_, err := c.client.
		Collection(ChildProfilePath).
		Doc(id.String()).
		Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
