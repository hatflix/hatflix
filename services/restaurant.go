//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=../test/restaurantservice.go
package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"easyfood/pkg/entity"
)

type RestaurantService interface {
	Get(ctx context.Context, id *int) ([]*entity.Restaurant, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*entity.Restaurant, error)
	GetByDish(ctx context.Context, dishID int) (*entity.Restaurant, error)
	Create(ctx context.Context, restaurant *entity.Restaurant) error
	Update(ctx context.Context, restaurant *entity.Restaurant) error
}

type restaurantService struct {
	db *sqlx.DB
}

func NewRestaurantService(db *sqlx.DB) RestaurantService {
	return restaurantService{db: db}
}

func (d restaurantService) Get(ctx context.Context, id *int) ([]*entity.Restaurant, error) {
	result := make([]*entity.Restaurant, 0)

	query := `
		SELECT * FROM restaurantes
	`

	if id != nil {
		query = fmt.Sprintf(`
			SELECT * FROM restaurantes WHERE id = %d
		`, *id)
	}

	err := d.db.SelectContext(ctx, &result, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (d restaurantService) GetByCategory(ctx context.Context, categoryID int) ([]*entity.Restaurant, error) {
	result := make([]*entity.Restaurant, 0)

	query := fmt.Sprintf(
		"SELECT r.id, r.horario_abertura, r.horario_fechamento, r.id_cidade,"+
			"r.dias_funcionamento, r.nome, r.descricao, r.telefone, r.endereco "+
			"FROM restaurantes r "+
			"INNER JOIN `restaurante-categoria` rc ON rc.id_restaurante = r.id "+
			"WHERE rc.id_categoria = %d", categoryID)

	err := d.db.SelectContext(ctx, &result, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (d restaurantService) GetByDish(ctx context.Context, dishID int) (*entity.Restaurant, error) {
	result := new(entity.Restaurant)

	query := `
		SELECT r.id, r.horario_abertura, r.horario_fechamento, r.id_cidade, 
		r.dias_funcionamento, r.nome, r.descricao, r.telefone, r.endereco 
		FROM restaurantes r INNER JOIN pratos p ON p.id_restaurante = r.id
		WHERE p.id = ?
	`

	err := d.db.GetContext(ctx, result, query, dishID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (d restaurantService) Create(ctx context.Context, restaurant *entity.Restaurant) error {
	query := `
		INSERT INTO restaurantes (horario_abertura, horario_fechamento, id_cidade, dias_funcionamento, nome, descricao, telefone, endereco)
		VALUES (:horario_abertura, :horario_fechamento, :id_cidade, :dias_funcionamento, :nome, :descricao, :telefone, :endereco)
	`

	result, err := d.db.NamedExecContext(ctx, query, restaurant)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	restaurant.Id = int(id)
	return nil
}

func (d restaurantService) Update(ctx context.Context, restaurant *entity.Restaurant) error {
	query := `
		UPDATE restaurantes
		SET horario_abertura = :horario_abertura, horario_fechamento = :horario_fechamento, 
			id_cidade = :id_cidade, dias_funcionamento = :dias_funcionamento, nome = :nome,
			descricao = :descricao, telefone = :telefone, endereco = :endereco
		WHERE id = :id
	`
	_, err := d.db.NamedExecContext(ctx, query, restaurant)
	if err != nil {
		return err
	}

	return nil
}
