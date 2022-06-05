package models

import "hatflix/pkg/entity"

func NewCategory(c ...*entity.Category) []*Category {
	result := make([]*Category, 0)
	for _, category := range c {
		if category.IsEmpty() {
			continue
		}
		result = append(result, &Category{
			ID:   category.Id,
			Name: category.Name,
		})
	}

	return result
}
