package graphql

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"easyfood/pkg/entity"
	"easyfood/services"
	mock "easyfood/test"
)

func TestQueryResolver_Category(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	categoryService := mock.NewMockCategoryService(ctrl)
	srvc := services.All{
		Category:   categoryService,
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
				Id: 201,
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

func TestQueryResolver_Dish(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	dishService := mock.NewMockDishService(ctrl)
	srvc := services.All{
		Dish:   dishService,
	}
	query := NewQueryResolver(srvc)

	t.Run("fail", func(t *testing.T) {
		Id := 500
		expectedErr := errors.New("failed")
		dishService.EXPECT().Get(gomock.Any(), &Id).Return(nil, expectedErr)
		res, err := query.Dish(ctx, &Id)

		require.Nil(t, res)
		require.True(t, errors.Is(expectedErr, err))
	})

	t.Run("success", func(t *testing.T) {
		Id := 500
		dish := []*entity.Dish{
			{
				Id:           200,
				Name:         "macarronada",
				Price:        15,
				CookTime:     25,
			},
		}
		dishService.EXPECT().Get(gomock.Any(), &Id).Return(dish, nil)
		res, err := query.Dish(ctx, &Id)

		require.Nil(t, err)
		require.Equal(t, "macarronada", res[0].Name)
	})
}
