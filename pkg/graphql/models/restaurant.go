package models

import (
	"hatflix/pkg/entity"
)

func NewStore(stores ...*entity.Store) []*Store {
	result := make([]*Store, 0)
	for _, store := range stores {
		if store.IsEmpty() {
			continue
		}
		result = append(result, &Store{
			ID:          store.Id,
			Name:        store.Name,
			Cnpj:        store.Cnpj,
			PhoneNumber: store.PhoneNumber,
			Address:     store.Address,
			Category:    store.CategoryID,
		})
	}

	return result
}
