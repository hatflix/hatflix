package graphql

import (
	"context"

	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
	"easyfood/services"
)

type restaurantResolver struct {
	services services.All
}

func NewRestaurantResolver(s services.All) gqlgen.RestaurantResolver {
	return restaurantResolver{services: s}
}

func (r restaurantResolver) Dishes(ctx context.Context, restaurant *models.Restaurant) ([]*models.Dish, error) {
	dishes, err := r.services.Dish.GetByRestaurant(ctx, restaurant.ID)
	if err != nil {
		return nil, err
	}

	return models.NewDish(dishes...), nil
}

func (r restaurantResolver) Category(ctx context.Context, restaurant *models.Restaurant) ([]*models.Category, error) {
	category, err := r.services.Category.GetByRestaurant(ctx, restaurant.ID)
	if err != nil {
		return nil, err
	}

	return models.NewCategory(category...), nil
}
