package model

import (
	"ref/entity"

	"gorm.io/gorm"
)

type Ingredient struct {
	ID         string
	Name       string
	Nutritions []*Nutrition `gorm:"many2many:ingredient_nutritions"`
	Weight     uint
	gorm.Model
}

func (m Ingredient) ToEntity() *entity.Ingredient {
	nutritions := []*entity.Nutrition{}
	if m.Nutritions != nil {
		for _, nutrition := range m.Nutritions {
			nutritions = append(nutritions, nutrition.ToEntity())
		}
	}
	return &entity.Ingredient{
		ID:         m.ID,
		Name:       m.Name,
		Nutritions: nutritions,
		Weight:     m.Weight,
	}
}
