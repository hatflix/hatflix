package graphql

import (
	"context"
	"errors"
	"hatflix/pkg/entity"
	"hatflix/pkg/graphql/gqlgen"
	"hatflix/pkg/graphql/models"
	"hatflix/services"
)

type mutationResolver struct {
	services services.All
}

func NewMutationResolver(s services.All) gqlgen.MutationResolver {
	return mutationResolver{services: s}
}

func (m mutationResolver) CreateClothes(ctx context.Context, input models.CreateClothInput) (*models.Clothes, error) {
	if input.Name == "" {
		return nil, errors.New("invalid dish name")
	}

	if input.IDStore == 0 {
		return nil, errors.New("must be associated to a store")
	}

	if input.Price == 0 {
		return nil, errors.New("must be a price")
	}

	if input.Size == "" {
		return nil, errors.New("must be a size")
	}

	if input.Quantity == 0 {
		return nil, errors.New("must be a quantity")
	}

	cloth := entity.Clothes{
		Name:       input.Name,
		StoreID:    input.IDStore,
		CategoryID: input.Category,
		Size:       input.Size,
		Price:      input.Price,
		Quantity:   input.Quantity,
	}

	err := m.services.Cloth.Create(ctx, &cloth)
	if err != nil {
		return nil, err
	}

	return models.NewCloth(&cloth)[0], nil
}

func (m mutationResolver) CreateCategory(ctx context.Context, input *models.CreateCategoryInput) (bool, error) {
	if input.Name == "" {
		return false, errors.New("must be a name")
	}

	category := entity.Category{
		Name: input.Name,
	}

	err := m.services.Category.Create(ctx, &category)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (m mutationResolver) CreateStore(ctx context.Context, input models.CreateStoreInput) (*models.Store, error) {

	if input.Name == "" {
		return nil, errors.New("must have a name")
	}
	if input.Cnpj == "" {
		return nil, errors.New("must have a Cnpj")
	}
	if input.Address == "" {
		return nil, errors.New("must have a Address")
	}
	if input.PhoneNumber == "" {
		return nil, errors.New("must have a PhoneNumber")
	}
	if input.Category == 0 {
		return nil, errors.New("must have a CategoryID")
	}

	store := entity.Store{
		Name:        input.Name,
		Cnpj:        input.Cnpj,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		CategoryID:  input.Category,
	}

	err := m.services.Store.Create(ctx, &store)
	if err != nil {
		return nil, err
	}

	return models.NewStore(&store)[0], nil

}

func (m mutationResolver) UpdateStore(ctx context.Context, input models.UpdateStoreInput) (*models.Store, error) {
	store, err := m.services.Store.Get(ctx, &input.ID)
	if err != nil {
		return nil, err
	}
	r := store[0]

	changes := false
	if &input.Name != nil {
		if input.Name != r.Name {
			r.Name = input.Name
			changes = true
		}
	}
	if &input.PhoneNumber != nil {
		if input.PhoneNumber != r.PhoneNumber {
			r.PhoneNumber = input.PhoneNumber
			changes = true
		}
	}
	if &input.Address != nil {
		if input.Address != r.Address {
			r.Address = input.Address
			changes = true
		}
	}
	if &input.Cnpj != nil {
		if input.Cnpj != r.Cnpj {
			r.Cnpj = input.Cnpj
			changes = true
		}
	}
	if &input.Category != nil {
		if input.Category != r.CategoryID {
			r.CategoryID = input.Category
			changes = true
		}
	}

	if changes {
		if err := m.services.Store.Update(ctx, r); err != nil {
			return nil, errors.New("failed to update" + err.Error())
		}
	}

	result := models.NewStore(r)[0]
	return result, nil
}

func (m mutationResolver) UpdateCloth(ctx context.Context, input models.UpdateClothInput) (*models.Clothes, error) {
	cloth, err := m.services.Cloth.Get(ctx, &input.ID)
	if err != nil {
		return nil, err
	}
	r := cloth[0]

	changes := false
	if &input.Name != nil {
		if input.Name != r.Name {
			r.Name = input.Name
			changes = true
		}
	}
	if &input.IDStore != nil {
		if input.IDStore != r.StoreID {
			r.StoreID = input.IDStore
			changes = true
		}
	}
	if &input.Size != nil {
		if input.Size != r.Size {
			r.Size = input.Size
			changes = true
		}
	}
	if &input.Price != nil {
		if input.Price != r.Price {
			r.Price = input.Price
			changes = true
		}
	}
	if &input.Quantity != nil {
		if input.Quantity != r.Quantity {
			r.Quantity = input.Quantity
			changes = true
		}
	}
	if &input.Category != nil {
		if input.Category != r.CategoryID {
			r.CategoryID = input.Category
			changes = true
		}
	}

	if changes {
		if err := m.services.Cloth.Update(ctx, r); err != nil {
			return nil, errors.New("failed to update" + err.Error())
		}
	}

	result := models.NewCloth(r)[0]
	return result, nil
}

func (m mutationResolver) UpdateCategory(ctx context.Context, input models.UpdateCategoryInput) (bool, error) {
	category := entity.Category{
		Id:   input.ID,
		Name: input.Name,
	}
	err := m.services.Category.Update(ctx, &category)
	if err != nil {
		return false, err
	}
	return true, nil
}
