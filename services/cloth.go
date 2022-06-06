//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=../test/dishservice.go
package services

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"hatflix/pkg/entity"
)

type ClothService interface {
	Get(ctx context.Context, id *int) ([]*entity.Clothes, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*entity.Clothes, error)
	GetByStore(ctx context.Context, restaurantID int) ([]*entity.Clothes, error)
	Create(ctx context.Context, Clothes *entity.Clothes) error
	Update(ctx context.Context, Clothes *entity.Clothes) error
}

type clothService struct {
	db *sqlx.DB
}

func NewClothService(db *sqlx.DB) ClothService {
	return clothService{db: db}
}

func (d clothService) Get(ctx context.Context, id *int) ([]*entity.Clothes, error) {
	result := make([]*entity.Clothes, 0)
	query := `
		SELECT * FROM produtos
	`
	if id != nil {
		query = fmt.Sprintf("SELECT * FROM produtos WHERE id = %d", *id)
	}

	err := d.db.Select(&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d clothService) GetByCategory(ctx context.Context, categoryID int) ([]*entity.Clothes, error) {
	result := make([]*entity.Clothes, 0)

	query := `
		SELECT * FROM produtos WHERE id_categoria = ?
	`

	err := d.db.Select(&result, query, categoryID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d clothService) GetByStore(ctx context.Context, idLoja int) ([]*entity.Clothes, error) {
	result := make([]*entity.Clothes, 0)

	query := `
		SELECT * FROM produtos WHERE id_loja = ?
	`

	err := d.db.Select(&result, query, idLoja)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d clothService) Create(ctx context.Context, Clothes *entity.Clothes) error {
	query := `
		INSERT INTO produtos (id_loja, id_categoria, nome, preco, quantidade, tamanho)
		VALUES (:id_loja, :id_categoria, :nome, :preco, :quantidade, :tamanho)
	`
	result, err := d.db.NamedExecContext(ctx, query, Clothes)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	Clothes.Id = int(id)
	return nil
}

func (d clothService) Update(ctx context.Context, Clothes *entity.Clothes) error {
	query := `
		UPDATE produtos
		SET id_categoria = :id_categoria, nome = :nome, preco = :preco, 
		    quantidade = :quantidade, tamanho = :tamanho
		WHERE id = :id
	`
	_, err := d.db.NamedExecContext(ctx, query, Clothes)
	if err != nil {
		return err
	}

	return nil
}
