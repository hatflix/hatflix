//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=../test/userservice.go
package services

import (
	"context"

	"github.com/jmoiron/sqlx"

	"easyfood/pkg/entity"
)

type UserService interface {
	Get(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type userService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return userService{db: db}
}

func (u userService) Get(ctx context.Context, id int) (*entity.User, error) {
	result := new(entity.User)
	query := "SELECT * FROM usuarios WHERE id = ?"

	if err := u.db.GetContext(ctx, &result, query, id); err != nil {
		return nil, err
	}

	return result, nil
}

func (u userService) Create(ctx context.Context, user *entity.User) error {
	query := `
	INSERT INTO usuarios (primeiro_nome, ultimo_nome, email, telefone, senha_hash)
	VALUES (:primeiro_nome, :ultimo_nome, :email, :telefone, SHA2(:senha_hash, 256))`

	result, err := u.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	user.Id = int(id)
	return nil
}
