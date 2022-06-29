package graphql

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"hatflix/pkg/entity"
	"hatflix/services"
	mock "hatflix/test"
)

func TestQueryResolver_Category(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	categoryService := mock.NewMockCategoryService(ctrl)
	srvc := services.All{
		Category: categoryService,
	}
	query := NewQueryResolver(srvc)

	t.Run("fail", func(t *testing.T) {
		id := 500
		expectedErr := errors.New("failed")
		categoryService.EXPECT().Get(gomock.Any(), &id).Return(nil, expectedErr)
		res, err := query.Category(ctx, &id)

		require.Nil(t, res)
		require.True(t, errors.Is(err, expectedErr))
	})

	t.Run("success all", func(t *testing.T) {
		expectedCategory := []*entity.Category{
			{
				Id:   200,
				Name: "asiática",
			},
			{
				Id:   201,
				Name: "brasileira",
			},
		}
		categoryService.EXPECT().Get(gomock.Any(), nil).Return(expectedCategory, nil)
		res, err := query.Category(ctx, nil)

		require.Nil(t, err)
		require.Equal(t, "brasileira", res[1].Name)
	})

	t.Run("success one", func(t *testing.T) {
		Id := 200
		expectedCategory := []*entity.Category{
			{
				Id:   200,
				Name: "asiática",
			},
		}
		categoryService.EXPECT().Get(gomock.Any(), &Id).Return(expectedCategory, nil)
		res, err := query.Category(ctx, &Id)

		require.Nil(t, err)
		require.Len(t, res, 1)
		require.Equal(t, "asiática", res[0].Name)
	})
}

func TestQueryResolver_Clothes(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	clothsService := mock.NewMockClothService(ctrl)
	srvc := services.All{
		Cloth: clothsService,
	}
	query := NewQueryResolver(srvc)

	t.Run("fail", func(t *testing.T) {
		id := 500
		expectedErr := errors.New("failed")
		clothsService.EXPECT().Get(gomock.Any(), &id).Return(nil, expectedErr)
		res, err := query.Clothes(ctx, &id)

		require.Nil(t, res)
		require.True(t, errors.Is(err, expectedErr))
	})

	t.Run("success all", func(t *testing.T) {
		expectedCategory := []*entity.Clothes{
			{
				Id:         1,
				Name:       "teste",
				StoreID:    1,
				CategoryID: 1,
				Size:       "M",
				Price:      10,
				Quantity:   1,
			},
			{
				Id:         1,
				Name:       "teste2",
				StoreID:    1,
				CategoryID: 1,
				Size:       "M",
				Price:      10,
				Quantity:   1,
			},
		}
		clothsService.EXPECT().Get(gomock.Any(), nil).Return(expectedCategory, nil)
		res, err := query.Clothes(ctx, nil)

		require.Nil(t, err)
		require.Equal(t, "teste2", res[1].Name)
	})

	t.Run("success one", func(t *testing.T) {
		Id := 200
		expectedCategory := []*entity.Clothes{
			{
				Id:         1,
				Name:       "teste",
				StoreID:    1,
				CategoryID: 1,
				Size:       "M",
				Price:      10,
				Quantity:   1,
			},
		}
		clothsService.EXPECT().Get(gomock.Any(), &Id).Return(expectedCategory, nil)
		res, err := query.Clothes(ctx, &Id)

		require.Nil(t, err)
		require.Len(t, res, 1)
		require.Equal(t, "teste", res[0].Name)
	})
}

func TestQueryResolver_Store(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	storesService := mock.NewMockStoreService(ctrl)
	srvc := services.All{
		Store: storesService,
	}
	query := NewQueryResolver(srvc)

	t.Run("fail", func(t *testing.T) {
		id := 500
		expectedErr := errors.New("failed")
		storesService.EXPECT().Get(gomock.Any(), &id).Return(nil, expectedErr)
		res, err := query.Store(ctx, &id)

		require.Nil(t, res)
		require.True(t, errors.Is(err, expectedErr))
	})

	t.Run("success all", func(t *testing.T) {
		expectedCategory := []*entity.Store{
			{
				Id:          1,
				Name:        "Nike",
				Cnpj:        "00000001",
				PhoneNumber: "31313131",
				Address:     "Rua torta",
				CategoryID:  1,
			},
			{
				Id:          2,
				Name:        "Adidas",
				Cnpj:        "00000002",
				PhoneNumber: "31313132",
				Address:     "Rua reta",
				CategoryID:  2,
			},
		}
		storesService.EXPECT().Get(gomock.Any(), nil).Return(expectedCategory, nil)
		res, err := query.Store(ctx, nil)

		require.Nil(t, err)
		require.Equal(t, "Adidas", res[1].Name)
	})

	t.Run("success one", func(t *testing.T) {
		Id := 200
		expectedCategory := []*entity.Store{
			{
				Id:          1,
				Name:        "Nike",
				Cnpj:        "00000001",
				PhoneNumber: "31313131",
				Address:     "Rua torta",
				CategoryID:  1,
			},
		}
		storesService.EXPECT().Get(gomock.Any(), &Id).Return(expectedCategory, nil)
		res, err := query.Store(ctx, &Id)

		require.Nil(t, err)
		require.Len(t, res, 1)
		require.Equal(t, "Nike", res[0].Name)
	})
}
