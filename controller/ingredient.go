package controller

import (
	"errors"
	"net/http"
	"ref/entity"

	"github.com/gin-gonic/gin"
)

type IngredientController struct {
	IngredientUseCase
}

func NewIngredientController(u IngredientUseCase) *IngredientController {
	return &IngredientController{u}
}

type GetIngredientsQuery struct {
	Min *uint `form:"min,omitempty"`
	Max *uint `form:"max,omitempty"`
}

type GetIngredientsResponse struct {
	Ingredients []*entity.Ingredient `json:"ingredients"`
}

// GetIngredients godoc
//
// @Summary	全素材取得 API
// @Description
// @Tags		Ingredient
// @Accept		json
// @Produce		json
// @Param		min		query		uint	false	"min"
// @Param		max		query		uint	false	"max"
// @Success		200		{object}	GetIngredientsResponse
// @Failure		401		{object}	entity.ErrorResponse
// @Failure		404		{object}	entity.ErrorResponse
// @Router		/ingredients	[get]
func (c IngredientController) GetIngredients(ctx *gin.Context) (interface{}, error) {
	var query GetIngredientsQuery
	err := ctx.ShouldBind(&query)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	if query.Min != nil && query.Max != nil {
		if *query.Min > *query.Max {
			return nil, entity.WrapError(http.StatusBadRequest, errors.New("Min greater than Max cannot be used."))
		}
	}
	ingredients, err := c.IngredientUseCase.GetIngredients(ctx, query.Min, query.Max)
	if err != nil {
		return nil, err
	}
	return GetIngredientsResponse{Ingredients: ingredients}, nil
}

type UpdateIngredientsQuery struct {
	Weight uint `form:"weight"`
}

// UpdateIngredients godoc
//
// @Summary	素材更新 API
// @Description
// @Tags		Ingredient
// @Accept		json
// @Produce		json
// @Param		ingredientId	path		string	true			"素材 ID"
// @Param		weight			query		uint	ture			"重さ"
// @Success		200				"OK"
// @Failure		401				{object}	entity.ErrorResponse
// @Failure		404				{object}	entity.ErrorResponse
// @Router		/ingredients/{ingredientId}	[put]
func (c IngredientController) UpdateIngredients(ctx *gin.Context) (interface{}, error) {
	var query UpdateIngredientsQuery
	err := ctx.ShouldBind(&query)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	ingredientId := ctx.Param("ingredientId")
	return nil, c.IngredientUseCase.UpdateIngredients(ctx, ingredientId, query.Weight)
}
