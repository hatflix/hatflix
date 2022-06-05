package graphql

import (
	"github.com/jmoiron/sqlx"

	"easyfood/pkg/graphql/gqlgen"
	"easyfood/services"
)

type app struct {
	services services.All
}

func NewResolverRoot(db *sqlx.DB) gqlgen.ResolverRoot {
	return app{services: services.NewServices(db)}
}

func (a app) Mutation() gqlgen.MutationResolver {
	return NewMutationResolver(a.services)
}

func (a app) Query() gqlgen.QueryResolver {
	return NewQueryResolver(a.services)
}

func (a app) Restaurant() gqlgen.RestaurantResolver {
	return NewRestaurantResolver(a.services)
}

func (a app) Dish() gqlgen.DishResolver {
	return NewDishResolver(a.services)
}

func (a app) Category() gqlgen.CategoryResolver {
	return NewCategoryResolver(a.services)
}
