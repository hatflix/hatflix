package graphql

import (
	"context"

	"hatflix/pkg/graphql/gqlgen"
	"hatflix/pkg/graphql/models"
	"hatflix/services"
)

type queryResolver struct {
	services services.All
}

func NewQueryResolver(services services.All) gqlgen.QueryResolver {
	return queryResolver{services: services}
}

func (q queryResolver) Category(ctx context.Context, id *int) ([]*models.Category, error) {
	categoryEntity, err := q.services.Category.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	category := models.NewCategory(categoryEntity...)
	return category, nil
}

func (q queryResolver) Clothes(ctx context.Context, id *int) ([]*models.Clothes, error) {
	d, err := q.services.Cloth.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.NewCloth(d...), nil
}

func (q queryResolver) Store(ctx context.Context, id *int) ([]*models.Store, error) {
	u, err := q.services.Store.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.NewStore(u...), nil
}
