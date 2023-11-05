package repository

import (
	"context"
	"fmt"
	"ref/entity"
	"ref/infrastructure/driver"
	"ref/infrastructure/repository/model"

	"gorm.io/gorm"
)

type IngredientRepository struct {
}

func NewIngredientRepository() *IngredientRepository {
	return &IngredientRepository{}
}

func (r IngredientRepository) GetIngredients(ctx context.Context, min, max *uint) ([]*entity.Ingredient, error) {
	records := []*model.Ingredient{}
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if min != nil {
		db = db.Where(fmt.Sprintf("weight >= %d", *min))
	}
	if max != nil {
		db = db.Where(fmt.Sprintf("weight <= %d", *max))
	}
	if err := db.Preload("Nutritions").Find(&records).Error; err != nil {
		return nil, err
	}
	var ingredients []*entity.Ingredient
	for _, record := range records {
		ingredients = append(ingredients, record.ToEntity())
	}
	return ingredients, nil
}
