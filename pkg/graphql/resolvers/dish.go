package graphql

import (
	"context"
	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
	"easyfood/services"
	"errors"
)

type dishResolver struct {
	services services.All
}

func NewDishResolver(s services.All) gqlgen.DishResolver {
	return dishResolver{services: s}
}

func (d dishResolver) Category(ctx context.Context, dish *models.Dish) (*models.Category, error) {
	categories, err := d.services.Category.GetByDish(ctx, dish.ID)
	if err != nil {
		return nil, err
	}
	if categories.IsEmpty() {
		return nil, nil
	}
	return models.NewCategory(categories)[0], nil
}

func (d dishResolver) Restaurant(ctx context.Context, dish *models.Dish) (*models.Restaurant, error) {
	restaurants, err := d.services.Restaurant.GetByDish(ctx, dish.ID)

	if err != nil {
		return nil, err
	}
	if restaurants.IsEmpty() {
		return &models.Restaurant{}, errors.New("Restaurant not found.")
	}
	return models.NewRestaurant(restaurants)[0], nil
}
