package services

import (
	"log"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/utils"
)

// GetUser - Returns user from db
func GetUser(username string) (models.User, error) {

	var user models.User

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	rows, err := dbConn.Query("SELECT * FROM app_user WHERE username = $1", username)
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Hash)
		if err != nil {
			log.Fatal(err)
			return user, err
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	return user, err

}
