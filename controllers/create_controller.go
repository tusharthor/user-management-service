package controllers

import (
	"copy_crud_api/service"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := service.CreateUser(db, ctx) //all we need to pass is db
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
	}
}
