package usecase

import (
	"context"
	"ref/entity"
)

type IngredientUseCase struct {
	IngredientRepository
}

func NewIngredientUseCase(repository IngredientRepository) *IngredientUseCase {
	return &IngredientUseCase{repository}
}

func (u IngredientUseCase) GetIngredients(ctx context.Context, min, max *uint) ([]*entity.Ingredient, error) {
	return u.IngredientRepository.GetIngredients(ctx, min, max)
}
