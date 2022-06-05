package graphql

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"easyfood/pkg/graphql/models"
	"easyfood/services"
	mock "easyfood/test"
)

func TestMutationResolver_UpdateCategory(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	categoryService := mock.NewMockCategoryService(ctrl)
	mutation := NewMutationResolver(services.All{
		Category: categoryService,
	})

	t.Run("fail", func(t *testing.T) {
		input := models.UpdateCategoryInput{
			ID:   500,
			Name: "failed",
		}
		expectedErr := errors.New("failed")
		categoryService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedErr)

		ok, err := mutation.UpdateCategory(ctx, input)
		require.False(t, ok)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		input := models.UpdateCategoryInput{
			ID:   200,
			Name: "success",
		}
		categoryService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

		ok, err := mutation.UpdateCategory(ctx, input)
		require.True(t, ok)
		require.Nil(t, err)
	})
}

func TestMutationResolver_UpdateDish(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	dishService := mock.NewMockDishService(ctrl)
	mutation := NewMutationResolver(services.All{
		Dish: dishService,
	})

	t.Run("fail", func(t *testing.T) {
		input := models.UpdateDishInput{
			ID:   500,
			Name: "failed",
		}
		expectedErr := errors.New("failed")
		dishService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedErr)

		dish, err := mutation.UpdateDish(ctx, input)
		require.Nil(t, dish)
		require.Error(t, err)
	})

	t.Run("invalid name", func(t *testing.T) {
		input := models.UpdateDishInput{
			ID:   500,
			Name: "",
		}

		dish, err := mutation.UpdateDish(ctx, input)
		require.Nil(t, dish)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		input := models.UpdateDishInput{
			ID:   200,
			Name: "success",
		}
		dishService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

		dish, err := mutation.UpdateDish(ctx, input)
		require.NotNil(t, dish)
		require.Equal(t, "success", dish.Name)
		require.Nil(t, err)
	})
}
