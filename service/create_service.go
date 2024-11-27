package service

import (
	"copy_crud_api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(db *sql.DB, ctx *gin.Context) error {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json data"})
		return err
	}

	_, err := db.Exec("insert into user (name,phone) values (?,?)", user.Name, user.Phone)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not create"})
		return err
	}
	// ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
	return nil
}
