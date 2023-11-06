package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"ref/entity"
	"ref/infrastructure/driver"
	"ref/infrastructure/repository/model"

	"github.com/go-sql-driver/mysql"
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

func (r IngredientRepository) UpdateIngredients(ctx context.Context, ingredientId string, weight uint) error {
	var record *model.Ingredient
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.First(&record, "id = ?", ingredientId).Error; err != nil {
		return err
	}
	record.Weight = weight
	if err := db.Save(record).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == driver.ErrDuplicateEntryNumber {
			return entity.WrapError(http.StatusConflict, err)
		}
		return err
	}
	return nil
}
