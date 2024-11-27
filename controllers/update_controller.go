package controllers

import (
	"copy_crud_api/service"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//call service
		err := service.UpdateByID(db, ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
			return
		}
	}
}
