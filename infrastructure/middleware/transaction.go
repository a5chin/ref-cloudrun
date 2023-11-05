package middleware

import (
	"net/http"
	"ref/infrastructure/driver"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tx := db.Begin()
		defer func() {
			if http.StatusBadRequest <= ctx.Writer.Status() {
				tx.Rollback()
				return
			}
			tx.Commit()
		}()
		ctx.Set(driver.TxKey, tx)
		ctx.Next()
	}
}
