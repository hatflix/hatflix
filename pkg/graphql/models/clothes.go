package models

import "hatflix/pkg/entity"

func NewCloth(clothes ...*entity.Clothes) []*Clothes {
	result := make([]*Clothes, 0)

	for _, cloth := range clothes {
		if cloth.IsEmpty() {
			continue
		}
		result = append(result, &Clothes{
			ID:       cloth.Id,
			Name:     cloth.Name,
			IDStore:  cloth.StoreID,
			Category: cloth.CategoryID,
			Size:     cloth.Size,
			Price:    cloth.Price,
			Quantity: cloth.Quantity,
		})
	}

	return result
}
