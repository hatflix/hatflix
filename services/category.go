//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=../test/categoryservice.go
package services

import (
	"context"
	"database/sql"
	"easyfood/pkg/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CategoryService interface {
	Get(ctx context.Context, id *int) ([]*entity.Category, error)
	GetByDish(ctx context.Context, dishId int) (*entity.Category, error)
	GetByRestaurant(ctx context.Context, restaurantId int) ([]*entity.Category, error)
	Create(ctx context.Context, category *entity.Category) error
	Update(ctx context.Context, category *entity.Category) error
}

type categoryService struct {
	db *sqlx.DB
}

func NewCategoryService(db *sqlx.DB) CategoryService {
	return categoryService{db: db}
}

func (c categoryService) Get(ctx context.Context, id *int) ([]*entity.Category, error) {
	result := make([]*entity.Category, 0)

	query := `SELECT * FROM categorias`

	if id != nil {
		query = fmt.Sprintf(`SELECT * FROM categorias WHERE id = %d`, *id)
	}
	err := c.db.SelectContext(ctx, &result, query)

	if err != nil {
		return nil, err
	}
	return result, err
}

func (c categoryService) GetByDish(ctx context.Context, dishId int) (*entity.Category, error) {
	result := new(entity.Category)

	query := `SELECT c.id, c.nome FROM categorias as c 
				INNER JOIN pratos as p
				ON c.id = p.id_categoria
				WHERE p.id = ?`

	err := c.db.GetContext(ctx, result, query, dishId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return result, err
}

func (c categoryService) GetByRestaurant(ctx context.Context, restaurantId int) ([]*entity.Category, error) {
	result := make([]*entity.Category, 0)

	query := fmt.Sprintf("SELECT c.id, c.nome FROM categorias as c "+
				"INNER JOIN `restaurante-categoria` as rc " +
				"on c.id = rc.id_categoria "+
				"WHERE rc.id_restaurante = %d", restaurantId)

	err := c.db.SelectContext(ctx, &result, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return result, err
}

func (c categoryService) Create(ctx context.Context, category *entity.Category) error {
	query := `
		INSERT INTO categorias (id, nome) VALUES (:id, :nome)
	`

	result, err := c.db.NamedExecContext(ctx, query, category)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	category.Id = int(id)
	return nil
}

func (c categoryService) Update(ctx context.Context, category *entity.Category) error {
	query := `UPDATE categorias SET nome = :nome WHERE id = :id`

	_, err := c.db.NamedExecContext(ctx, query, category)
	if err != nil {
		return err
	}

	return nil
}
