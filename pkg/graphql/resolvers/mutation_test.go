package graphql

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"hatflix/pkg/graphql/models"
	"hatflix/services"
	mock "hatflix/test"
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

func TestMutationResolver_CreateCategory(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	categoryService := mock.NewMockCategoryService(ctrl)
	mutation := NewMutationResolver(services.All{
		Category: categoryService,
	})

	t.Run("fail", func(t *testing.T) {
		input := models.CreateCategoryInput{
			Name: "teste",
		}
		expectedErr := errors.New("failed")
		categoryService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(expectedErr)

		ok, err := mutation.CreateCategory(ctx, &input)
		require.False(t, ok)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		input := models.CreateCategoryInput{
			Name: "teste",
		}
		categoryService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

		ok, err := mutation.CreateCategory(ctx, &input)
		require.True(t, ok)
		require.Nil(t, err)
	})
}

//func TestMutationResolver_UpdateCloth(t *testing.T) {
//	ctx := context.Background()
//	ctrl := gomock.NewController(t)
//	clothService := mock.NewMockClothService(ctrl)
//	mutation := NewMutationResolver(services.All{
//		Cloth: clothService,
//	})
//
//	t.Run("fail", func(t *testing.T) {
//		input := models.UpdateClothInput{
//			ID:       0,
//			Name:     "teste",
//			IDStore:  0,
//			Category: 0,
//			Size:     "M",
//			Price:    0,
//			Quantity: 0,
//		}
//		expectedErr := errors.New("failed")
//		clothService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedErr)
//
//		_, err := mutation.UpdateCloth(ctx, input)
//		//require.Equal(t, "teste", ok.Name)
//		require.Error(t, err)
//	})
//
//	t.Run("success", func(t *testing.T) {
//		input := models.UpdateClothInput{
//			ID:       0,
//			Name:     "teste",
//			IDStore:  0,
//			Category: 0,
//			Size:     "M",
//			Price:    0,
//			Quantity: 0,
//		}
//		clothService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
//
//		res, err := mutation.UpdateCloth(ctx, input)
//		require.Equal(t, "teste", res.Name)
//		require.Nil(t, err)
//	})
//}
//
//func TestMutationResolver_CreateClothes(t *testing.T) {
//	ctx := context.Background()
//	ctrl := gomock.NewController(t)
//	clothService := mock.NewMockClothService(ctrl)
//	mutation := NewMutationResolver(services.All{
//		Cloth: clothService,
//	})
//
//	t.Run("fail", func(t *testing.T) {
//		input := models.CreateClothInput{
//			Name:     "camisa",
//			IDStore:  0,
//			Category: 0,
//			Size:     "M",
//			Price:    0,
//			Quantity: 0,
//		}
//		expectedErr := errors.New("failed")
//		clothService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(expectedErr)
//
//		ok, err := mutation.CreateClothes(ctx, input)
//		require.Equal(t, "camisa", ok.Name)
//		require.Error(t, err)
//	})
//
//	t.Run("success", func(t *testing.T) {
//		input := models.CreateClothInput{
//			Name:     "camisa",
//			IDStore:  0,
//			Category: 0,
//			Size:     "M",
//			Price:    0,
//			Quantity: 0,
//		}
//		clothService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
//
//		ok, err := mutation.CreateClothes(ctx, input)
//		require.Equal(t, "camisa", ok.Name)
//		require.Nil(t, err)
//	})
//}
