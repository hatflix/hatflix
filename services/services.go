package services

import "github.com/jmoiron/sqlx"

type All struct {
	User       UserService
	Dish       DishService
	Category   CategoryService
	Restaurant RestaurantService
}

func NewServices(db *sqlx.DB) All {
	return All{
		User:       NewUserService(db),
		Dish:       NewDishService(db),
		Category:   NewCategoryService(db),
		Restaurant: NewRestaurantService(db),
	}
}
