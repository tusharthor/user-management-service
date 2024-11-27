package service

import (
	"copy_crud_api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetByID(db *sql.DB, ctx *gin.Context) (models.User, error) {
	var user models.User
	// get id
	id := ctx.Param("id")
	result := db.QueryRow("select id,name,phone from user where id=?", id)
	err := result.Scan(&user.ID, &user.Name, &user.Phone)

	switch {
	case err == sql.ErrNoRows:
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return user, nil
	case err != nil: //means having error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return user, err
	default:
		ctx.JSON(http.StatusOK, user)
		return user, nil
	}
}
