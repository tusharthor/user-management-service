package controllers

import (
	"copy_crud_api/service"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllUsers(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := service.ListAll(db) //will get users
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//success then return all users
		ctx.JSON(http.StatusOK, gin.H{"users": users})
	}
}
