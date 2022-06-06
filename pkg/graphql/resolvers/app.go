package graphql

import (
	"github.com/jmoiron/sqlx"

	"hatflix/pkg/graphql/gqlgen"
	"hatflix/services"
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

//func (a app) Restaurant() gqlgen.StoreResolver {
//	return NewRestaurantResolver(a.services)
//}
//
//func (a app) Dish() gqlgen.ClothesResolver {
//	return NewDishResolver(a.services)
//}
//
//func (a app) Category() gqlgen.CategoryResolver {
//	return NewCategoryResolver(a.services)
//}
