package main

import (
	"fmt"
	"net/http"
	"ref/config"
	"ref/controller"
	"ref/entity"
	"ref/infrastructure/driver"
	"ref/infrastructure/middleware"
	"ref/infrastructure/repository"
	"ref/usecase"

	"github.com/gin-gonic/gin"
)

// @title		ref
// @description	ref
// @version		1.0
func main() {
	conf := config.Load()
	db := driver.NewDB(conf)

	ingredientRepository := repository.NewIngredientRepository()
	nutritionRepository := repository.NewNutritionRepository()

	ingredientUseCase := usecase.NewIngredientUseCase(ingredientRepository)
	nutritionUseCase := usecase.NewNutritionUseCase(nutritionRepository)

	ingredientController := controller.NewIngredientController(ingredientUseCase)
	nutritionController := controller.NewNutritionController(nutritionUseCase)

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction(db))

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "It works")
	})

	api := app.Group("/api/v1")

	ingredientRouter := api.Group("/ingredients")
	ingredientRouter.GET("/", handleResponse(ingredientController.GetIngredients))

	nutritionRouter := api.Group("/nutritions")
	nutritionRouter.GET("/", handleResponse(nutritionController.GetNutritions))
	nutritionRouter.GET("/:nutritionId", handleResponse(nutritionController.GetNutritionByID))

	runApp(app, conf)
}

func runApp(app *gin.Engine, conf *config.Config) {
	app.Run(
		fmt.Sprintf("%s:%s", conf.HOSTNAME, conf.PORT),
	)
}

func handleResponse(f func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := f(c)
		if err != nil {
			e, ok := err.(*entity.Error)
			if ok {
				c.JSON(e.Code, entity.ErrorResponse{Message: err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
			}
			c.Abort()
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
