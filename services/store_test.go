package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hatflix/pkg/entity"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestStoreService_Get(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewStoreService(sqlxDB)

	StoreId := 10

	query := regexp.QuoteMeta(fmt.Sprintf("SELECT * FROM lojas WHERE id = %d", StoreId))

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(query).
			WillReturnRows(sqlmock.NewRows([]string{"id", "nome"}).
				AddRow(StoreId, "massas"))

		res, err := srvc.Get(ctx, &StoreId)
		require.NoError(t, err)
		require.Equal(t, StoreId, res[0].Id)
		require.Equal(t, "massas", res[0].Name)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(StoreId).
			WillReturnError(expectedErr)

		res, err := srvc.Get(ctx, &StoreId)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestStoreService_GetByCategory(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewStoreService(sqlxDB)

	query := regexp.QuoteMeta(
		"SELECT r.id, r.horario_abertura, r.horario_fechamento, r.id_cidade,r.dias_funcionamento, r.nome, r.descricao, r.telefone, r.endereco " +
			"FROM lojas r INNER JOIN `categorias` rc ON rc.id_restaurante = r.id " +
			"WHERE rc.nome = 200")

	t.Run("success", func(t *testing.T) {
		dishId := 200
		mock.ExpectQuery(query).
			WithArgs(dishId)

		_, _ = srvc.GetByCategory(ctx, dishId)
		//require.NoError(t, err)
		//require.Equal(t, StoreId, res)
		//require.Equal(t, "massas", res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")
		dishId := 200

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		res, err := srvc.GetByCategory(ctx, dishId)
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("failed: no rows", func(t *testing.T) {
		expectedErr := sql.ErrNoRows
		dishId := 10
		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		res, err := srvc.GetByCategory(ctx, dishId)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestStoreService_Create(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewStoreService(sqlxDB)

	query := regexp.QuoteMeta(
		"INSERT INTO lojas (nome, cnpj, telefone, endereco, id_categoria) VALUES (?, ?, ?, ?, ?)")

	Store := entity.Store{
		Id:          1,
		Name:        "Nike",
		Cnpj:        "00000001",
		PhoneNumber: "31313131",
		Address:     "rua torta",
		CategoryID:  0,
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(Store.Name, Store.Cnpj, Store.PhoneNumber, Store.Address, Store.CategoryID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := srvc.Create(ctx, &Store)
		require.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(Store).
			WillReturnError(expectedErr)

		err := srvc.Create(ctx, &Store)
		require.Error(t, err)
	})
}

func TestStoreService_Update(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewStoreService(sqlxDB)

	query := regexp.QuoteMeta(
		"UPDATE produtos SET id_categoria = ?, nome = ?, preco = ?, quantidade = ?, tamanho = ? WHERE id = ?")

	Store := entity.Store{
		Id:          1,
		Name:        "Nike",
		Cnpj:        "00000001",
		PhoneNumber: "31313131",
		Address:     "rua torta",
		CategoryID:  0,
	}

	//t.Run("success", func(t *testing.T) {
	//	mock.ExpectExec(query).
	//		WithArgs(
	//			Store.CategoryID,
	//			Store.Name,
	//			Store.Price,
	//			Store.Quantity,
	//			Store.Size,
	//			Store.Id,
	//		)
	//
	//	err := srvc.Update(ctx, &Store)
	//	require.NoError(t, err)
	//	require.Equal(t, StoreId, res)
	//	require.Equal(t, "massas", res)
	//})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(Store).
			WillReturnError(expectedErr)

		err := srvc.Update(ctx, &Store)
		require.Error(t, err)
	})
}
