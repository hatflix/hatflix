package services

import "github.com/jmoiron/sqlx"

type All struct {
	Cloth    ClothService
	Category CategoryService
	Store    StoreService
}

func NewServices(db *sqlx.DB) All {
	return All{
		Cloth:    NewClothService(db),
		Category: NewCategoryService(db),
		Store:    NewStoreService(db),
	}
}
