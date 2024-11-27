package routes

import (
	"copy_crud_api/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	//get all users
	router.GET("/users", controllers.ListAllUsers(db))
	//create user
	router.POST("/createuser", controllers.CreateUser(db))
	//get user by id
	router.GET("/user/:id", controllers.GetUserById(db))
	//update user
	router.PUT("/user/:id", controllers.UpdateUser(db))
}
