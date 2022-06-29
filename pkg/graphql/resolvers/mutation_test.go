package graphql

import (
	"context"
	"errors"
	"hatflix/pkg/entity"
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

func TestMutationResolver_UpdateCloth(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	clothService := mock.NewMockClothService(ctrl)
	mutation := NewMutationResolver(services.All{
		Cloth: clothService,
	})

	mockCloth := []*entity.Clothes{{
		Id:         500,
		Name:       "teste",
		StoreID:    0,
		CategoryID: 0,
		Size:       "M",
		Price:      0,
		Quantity:   0,
	}}

	t.Run("fail update", func(t *testing.T) {
		input := models.UpdateClothInput{
			ID: 500,
		}
		clothService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mockCloth, nil)

		expectedErr := errors.New("failed")
		clothService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedErr)

		ok, err := mutation.UpdateCloth(ctx, input)
		require.Nil(t, ok)
		require.Error(t, err)
	})

	t.Run("fail get", func(t *testing.T) {
		input := models.UpdateClothInput{
			ID: 500,
		}

		expectedErr := errors.New("failed")
		clothService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, expectedErr)

		ok, err := mutation.UpdateCloth(ctx, input)
		require.Nil(t, ok)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		input := models.UpdateClothInput{
			ID:       0,
			Name:     "teste2",
			IDStore:  1,
			Category: 1,
			Size:     "GG",
			Price:    100,
			Quantity: 10,
		}
		clothService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mockCloth, nil)
		clothService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

		res, err := mutation.UpdateCloth(ctx, input)
		require.Equal(t, "teste2", res.Name)
		require.Nil(t, err)
	})
}

func TestMutationResolver_UpdateStore(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	storeService := mock.NewMockStoreService(ctrl)
	mutation := NewMutationResolver(services.All{
		Store: storeService,
	})

	mockStore := []*entity.Store{{
		Id:          500,
		Name:        "teste",
		Cnpj:        "00001",
		PhoneNumber: "31313131",
		Address:     "rua torta",
		CategoryID:  1,
	},
	}

	t.Run("fail update", func(t *testing.T) {
		input := models.UpdateStoreInput{
			ID: 500,
		}
		storeService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mockStore, nil)

		expectedErr := errors.New("failed")
		storeService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedErr)

		ok, err := mutation.UpdateStore(ctx, input)
		require.Nil(t, ok)
		require.Error(t, err)
	})

	t.Run("fail get", func(t *testing.T) {
		input := models.UpdateStoreInput{
			ID: 500,
		}

		expectedErr := errors.New("failed")
		storeService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, expectedErr)

		ok, err := mutation.UpdateStore(ctx, input)
		require.Nil(t, ok)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		input := models.UpdateStoreInput{
			ID:          500,
			Name:        "teste2",
			Cnpj:        "0002",
			PhoneNumber: "31313132",
			Address:     "rua reta",
			Category:    2,
		}
		storeService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mockStore, nil)
		storeService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

		res, err := mutation.UpdateStore(ctx, input)
		require.Equal(t, "teste2", res.Name)
		require.Nil(t, err)
	})
}

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
