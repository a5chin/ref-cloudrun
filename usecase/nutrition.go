package usecase

import (
	"context"
	"ref/entity"
)

type NutritionUseCase struct {
	NutritionRepository
}

func NewNutritionUseCase(repository NutritionRepository) *NutritionUseCase {
	return &NutritionUseCase{repository}
}

func (u NutritionUseCase) GetNutritions(ctx context.Context) ([]*entity.Nutrition, error) {
	return u.NutritionRepository.GetNutritions(ctx)
}

func (u NutritionUseCase) GetNutritionByID(ctx context.Context, nutritionId string) (*entity.Nutrition, error) {
	return u.NutritionRepository.GetNutritionByID(ctx, nutritionId)
}
