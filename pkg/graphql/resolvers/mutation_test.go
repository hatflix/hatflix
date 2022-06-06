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
