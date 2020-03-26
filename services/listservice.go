package services

import (
	"log"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/utils"
)

func GetList(listId int, userId int32) (models.List, error) {
	var list models.List

	query := "SELECT id, title, app_user_id " +
		"FROM list " +
		"WHERE id = $1"

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	row := dbConn.QueryRow(query, listId)

	err := row.Scan(&list.Id, &list.Title, &list.UserId)

	if err != nil {
		return list, &models.NotFoundError{}
	}

	if list.UserId != userId {
		return models.List{}, &models.AuthenticationError{}
	}

	return list, err
}

func GetLists(userId int32) ([]models.List, error) {
	var lists []models.List

	query := "SELECT id, title, app_user_id " +
		"FROM list " +
		"WHERE app_user_id = $1"

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	rows, err := dbConn.Query(query, userId)
	if err != nil {
		log.Fatal(err)
		return lists, err
	}

	for rows.Next() {
		list := models.List{}
		err := rows.Scan(&list.Id, &list.Title, &list.UserId)
		if err != nil {
			log.Fatal(err)
			return lists, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func SaveList(title string, userId int32) (models.List, error) {

	var list models.List

	query := "INSERT into list (title, app_user_id) " +
		"VALUES ($1, $2) " +
		"RETURNING title, id, app_user_id"

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	row := dbConn.QueryRow(query, title, userId)

	err := row.Scan(&list.Title, &list.Id, &list.UserId)

	if err != nil {
		return list, err
	}

	return list, nil
}
