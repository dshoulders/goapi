package services

import (
	"log"

	"github.com/dshoulders/goapi/utils"
)

// User - User details
type User struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"-"`
}

// GetUser - Returns user from db
func GetUser(username string) (User, error) {

	var user User

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
