package controllers

import (
	"copy_crud_api/service"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserById(db *sql.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		_, err := service.GetByID(db, ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
			return
		}

		// ctx.JSON(http.StatusOK, user)

	}
}
