package service

import (
	"copy_crud_api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateByID(db *sql.DB, ctx *gin.Context) error {
	var user models.User
	id := ctx.Param("id")

	//bind data
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return err
	}

	result := db.QueryRow("select id from user where id=?", id)
	er := result.Scan(&user.ID)

	if er == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return er
	}

	_, err := db.Exec("update user set name=?, phone=? where id=?", user.Name, user.Phone, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
	return nil
}
