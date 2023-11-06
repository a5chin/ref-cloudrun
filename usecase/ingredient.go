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

func (u IngredientUseCase) UpdateIngredients(ctx context.Context, ingredientId string, weight uint) error {
	return u.IngredientRepository.UpdateIngredients(ctx, ingredientId, weight)
}
