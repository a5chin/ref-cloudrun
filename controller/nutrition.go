package controller

import (
	"net/http"
	"ref/entity"

	"github.com/gin-gonic/gin"
)

type NutritionController struct {
	NutritionUseCase
}

func NewNutritionController(u NutritionUseCase) *NutritionController {
	return &NutritionController{u}
}

type GetNutritionsResponse struct {
	Nutritions []*entity.Nutrition `json:"nutritions"`
}

// GetNutritions godoc
//
// @Summary	全栄養素取得 API
// @Description
// @Tags		Nutrition
// @Accept		json
// @Produce		json
// @Success		200		{object}	GetNutritionsResponse	"OK"
// @Failure		401		{object}	entity.ErrorResponse
// @Failure		404		{object}	entity.ErrorResponse
// @Router		/nutritions	[get]
func (c NutritionController) GetNutritions(ctx *gin.Context) (interface{}, error) {
	nutritions, err := c.NutritionUseCase.GetNutritions(ctx)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetNutritionsResponse{Nutritions: nutritions}, nil
}

type GetNutritionResponse struct {
	Nutrition *entity.Nutrition `json:"nutrition"`
}

// GetNutritionByID godoc
//
// @Summary	栄養素取得 API
// @Description
// @Tags		Nutrition
// @Accept		json
// @Param		nutritionId	path		string	true			"栄養素 ID"
// @Produce		json
// @Success		200			{object}	GetNutritionsResponse	"OK"
// @Failure		401			{object}	entity.ErrorResponse
// @Failure		404			{object}	entity.ErrorResponse
// @Router		/nutritions/{nutritionId}	[get]
func (c NutritionController) GetNutritionByID(ctx *gin.Context) (interface{}, error) {
	nutritionId := ctx.Param("nutritionId")
	nutrition, err := c.NutritionUseCase.GetNutritionByID(ctx, nutritionId)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetNutritionResponse{Nutrition: nutrition}, nil
}
