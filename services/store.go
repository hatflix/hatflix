//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=../test/restaurantservice.go
package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"hatflix/pkg/entity"
)

type StoreService interface {
	Get(ctx context.Context, id *int) ([]*entity.Store, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*entity.Store, error)
	Create(ctx context.Context, Store *entity.Store) error
	Update(ctx context.Context, Store *entity.Store) error
}

type storeService struct {
	db *sqlx.DB
}

func NewStoreService(db *sqlx.DB) StoreService {
	return storeService{db: db}
}

func (d storeService) Get(ctx context.Context, id *int) ([]*entity.Store, error) {
	result := make([]*entity.Store, 0)

	query := `
		SELECT * FROM lojas
	`

	if id != nil {
		query = fmt.Sprintf(`
			SELECT * FROM lojas WHERE id = %d
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

func (d storeService) GetByCategory(ctx context.Context, categoryName int) ([]*entity.Store, error) {
	result := make([]*entity.Store, 0)

	query := fmt.Sprintf(
		"SELECT r.id, r.horario_abertura, r.horario_fechamento, r.id_cidade,"+
			"r.dias_funcionamento, r.nome, r.descricao, r.telefone, r.endereco "+
			"FROM lojas r "+
			"INNER JOIN `categorias` rc ON rc.id_restaurante = r.id "+
			"WHERE rc.nome = %d", categoryName)

	err := d.db.SelectContext(ctx, &result, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (d storeService) Create(ctx context.Context, Store *entity.Store) error {
	query := `
		INSERT INTO lojas (nome, cnpj, telefone, endereco, id_categoria)
		VALUES (:nome, :cnpj, :telefone, :endereco, :id_categoria)
	`

	result, err := d.db.NamedExecContext(ctx, query, Store)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	Store.Id = int(id)
	return nil
}

func (d storeService) Update(ctx context.Context, Store *entity.Store) error {
	query := `
		UPDATE lojas
		SET nome = :nome, cnpj = :cnpj, telefone = :telefone, endereco = :endereco, id_categoria = :id_categoria
		WHERE id = :id
	`
	_, err := d.db.NamedExecContext(ctx, query, Store)
	if err != nil {
		return err
	}

	return nil
}
