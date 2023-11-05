package model

import (
	"ref/entity"
)

type Nutrition struct {
	ID   string
	Name string
}

func (m Nutrition) ToEntity() *entity.Nutrition {
	return &entity.Nutrition{
		ID:   m.ID,
		Name: m.Name,
	}
}
