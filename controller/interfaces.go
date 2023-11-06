package controller

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock.go -package=$GOPACKAGE

import (
	"context"
	"ref/entity"
)

type IngredientUseCase interface {
	GetIngredients(ctx context.Context, min, max *uint) ([]*entity.Ingredient, error)
	UpdateIngredients(ctx context.Context, ingredientId string, weight uint) error
}

type NutritionUseCase interface {
	GetNutritions(ctx context.Context) ([]*entity.Nutrition, error)
	GetNutritionByID(ctx context.Context, nutritionId string) (*entity.Nutrition, error)
}
