//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=../test/dishservice.go
package services

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"easyfood/pkg/entity"
)

type DishService interface {
	Get(ctx context.Context, id *int) ([]*entity.Dish, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*entity.Dish, error)
	GetByRestaurant(ctx context.Context, restaurantID int) ([]*entity.Dish, error)
	Create(ctx context.Context, dish *entity.Dish) error
	Update(ctx context.Context, dish *entity.Dish) error
}

type dishService struct {
	db *sqlx.DB
}

func NewDishService(db *sqlx.DB) DishService {
	return dishService{db: db}
}

func (d dishService) Get(ctx context.Context, id *int) ([]*entity.Dish, error) {
	result := make([]*entity.Dish, 0)
	query := `
		SELECT * FROM pratos
	`
	if id != nil {
		query = fmt.Sprintf("SELECT * FROM pratos WHERE id = %d", *id)
	}

	err := d.db.Select(&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d dishService) GetByCategory(ctx context.Context, categoryID int) ([]*entity.Dish, error) {
	result := make([]*entity.Dish, 0)

	query := `
		SELECT * FROM pratos WHERE id_categoria = ?
	`

	err := d.db.Select(&result, query, categoryID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d dishService) GetByRestaurant(ctx context.Context, restaurantID int) ([]*entity.Dish, error) {
	result := make([]*entity.Dish, 0)

	query := `
		SELECT * FROM pratos WHERE id_restaurante = ?
	`

	err := d.db.Select(&result, query, restaurantID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d dishService) Create(ctx context.Context, dish *entity.Dish) error {
	query := `
		INSERT INTO pratos (id_restaurante, id_categoria, nome, preco, tempo_de_preparo)
		VALUES (:id_restaurante, :id_categoria, :nome, :preco, :tempo_de_preparo)
	`
	result, err := d.db.NamedExecContext(ctx, query, dish)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	dish.Id = int(id)
	return nil
}

func (d dishService) Update(ctx context.Context, dish *entity.Dish) error {
	query := `
		UPDATE pratos
		SET id_categoria = :id_categoria, nome = :nome, preco = :preco, 
		    tempo_de_preparo = :tempo_de_preparo
		WHERE id = :id
	`
	_, err := d.db.NamedExecContext(ctx, query, dish)
	if err != nil {
		return err
	}

	return nil
}
