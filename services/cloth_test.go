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

func TestClothService_Get(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewClothService(sqlxDB)

	ClothId := 10

	query := regexp.QuoteMeta(fmt.Sprintf("SELECT * FROM produtos WHERE id = %d", ClothId))

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(query).
			WillReturnRows(sqlmock.NewRows([]string{"id", "nome"}).
				AddRow(ClothId, "massas"))

		res, err := srvc.Get(ctx, &ClothId)
		require.NoError(t, err)
		require.Equal(t, ClothId, res[0].Id)
		require.Equal(t, "massas", res[0].Name)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(ClothId).
			WillReturnError(expectedErr)

		res, err := srvc.Get(ctx, &ClothId)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestClothService_GetByCategory(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewClothService(sqlxDB)

	query := regexp.QuoteMeta(
		"SELECT * FROM produtos WHERE id_categoria = ?")

	t.Run("success", func(t *testing.T) {
		dishId := 200
		ClothId := 10
		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "nome"}).
				AddRow(ClothId, "massas"))

		_, err := srvc.GetByCategory(ctx, dishId)
		require.NoError(t, err)
		//require.Equal(t, ClothId, res)
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

func TestClothService_GetByStore(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewClothService(sqlxDB)

	query := regexp.QuoteMeta(
		"SELECT * FROM produtos WHERE id_loja = ?")

	t.Run("success", func(t *testing.T) {
		dishId := 200
		ClothId := 10
		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "nome"}).
				AddRow(ClothId, "massas"))

		_, err := srvc.GetByStore(ctx, dishId)
		require.NoError(t, err)
		//require.Equal(t, ClothId, res)
		//require.Equal(t, "massas", res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")
		dishId := 200

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		res, err := srvc.GetByStore(ctx, dishId)
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("failed: no rows", func(t *testing.T) {
		expectedErr := sql.ErrNoRows
		dishId := 10
		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		res, err := srvc.GetByStore(ctx, dishId)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestClothService_Create(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewClothService(sqlxDB)

	query := regexp.QuoteMeta(
		"INSERT INTO produtos (id_loja, id_categoria, nome, preco, quantidade, tamanho) VALUES (?, ?, ?, ?, ?, ?)")

	cloth := entity.Clothes{
		Id:         1,
		Name:       "camisa",
		StoreID:    1,
		CategoryID: 1,
		Size:       "M",
		Price:      1,
		Quantity:   1,
	}

	//t.Run("success", func(t *testing.T) {
	//	mock.ExpectExec(query).
	//		WithArgs(
	//			cloth.StoreID,
	//			cloth.CategoryID,
	//			cloth.Name,
	//			cloth.Price,
	//			cloth.Quantity,
	//			cloth.Size,
	//		)
	//
	//	err := srvc.Create(ctx, &cloth)
	//	require.NoError(t, err)
	//require.Equal(t, ClothId, res)
	//require.Equal(t, "massas", res)
	//})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(cloth).
			WillReturnError(expectedErr)

		err := srvc.Create(ctx, &cloth)
		require.Error(t, err)
	})
}

func TestClothService_Update(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewClothService(sqlxDB)

	query := regexp.QuoteMeta(
		"UPDATE produtos SET id_categoria = ?, nome = ?, preco = ?, quantidade = ?, tamanho = ? WHERE id = ?")

	cloth := entity.Clothes{
		Id:         1,
		Name:       "camisa",
		StoreID:    1,
		CategoryID: 1,
		Size:       "M",
		Price:      1,
		Quantity:   1,
	}

	//t.Run("success", func(t *testing.T) {
	//	mock.ExpectExec(query).
	//		WithArgs(
	//			cloth.CategoryID,
	//			cloth.Name,
	//			cloth.Price,
	//			cloth.Quantity,
	//			cloth.Size,
	//			cloth.Id,
	//		)
	//
	//	err := srvc.Update(ctx, &cloth)
	//	require.NoError(t, err)
	//	require.Equal(t, ClothId, res)
	//	require.Equal(t, "massas", res)
	//})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(cloth).
			WillReturnError(expectedErr)

		err := srvc.Update(ctx, &cloth)
		require.Error(t, err)
	})
}
