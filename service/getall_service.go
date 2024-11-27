package service

import (
	"copy_crud_api/models"
	"database/sql"
)

// return slice of users
func ListAll(db *sql.DB) ([]models.User, error) {
	query := "select id,name,phone from user"
	rows, err := db.Query(query) //return rows
	if err != nil {
		return nil, err
	}
	//close before exiting function
	defer rows.Close()

	//create slice
	var userList []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Phone)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}

	//finally return users
	return userList, nil
}
