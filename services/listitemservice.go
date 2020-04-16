package services

import (
	"log"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/utils"
)

func getListOwner(listId int) (int, error) {
	query := "SELECT app_user_id" +
		" FROM list" +
		" WHERE id = $1"

	var ownerId int

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	row := dbConn.QueryRow(query, listId)

	err := row.Scan(&ownerId)

	if err != nil {
		return 0, err
	}

	return ownerId, nil
}

func GetListItems(listId int, userId int32) ([]models.ListItem, error) {
	var listItems []models.ListItem

	ownerId, err := getListOwner(listId)

	if err != nil {
		return listItems, &models.NotFoundError{}
	}

	if int(userId) != ownerId {
		return listItems, &models.AuthenticationError{}
	}

	query := "SELECT li.title, li.id, li.notes, au.id" +
		" FROM list_item li" +
		" INNER JOIN app_user au" +
		"    on li.app_user_id = au.id" +
		" WHERE list_id = $1"

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	rows, err := dbConn.Query(query, listId)
	if err != nil {
		log.Fatal(err)
		return listItems, err
	}

	for rows.Next() {
		listItem := models.ListItem{}
		err := rows.Scan(&listItem.Title, &listItem.Id, &listItem.Notes, &listItem.UserId)
		if err != nil {
			log.Fatal(err)
			return listItems, err
		}
		listItems = append(listItems, listItem)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return listItems, err
	}

	return listItems, nil
}

func SaveListItem(listId int, listItem models.ListItem) (models.ListItem, error) {

	query := `
		INSERT into list_item (title, notes, list_id, app_user_id)
		SELECT $1, $2, $3, $4
		WHERE EXISTS (
			SELECT * FROM list 
			WHERE id = $3 AND app_user_id = $4
		)
		RETURNING id
		`

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	row := dbConn.QueryRow(query, listItem.Title, listItem.Notes, listId, listItem.UserId)

	err := row.Scan(&listItem.Id)

	if err != nil {
		return listItem, err
	}

	return listItem, nil
}

func GetListItem(listItemId int, userId int32) (models.ListItem, error) {

	var listItem models.ListItem

	query := "SELECT title, id, notes, app_user_id" +
		" FROM list_item" +
		" WHERE id = $1"

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	row := dbConn.QueryRow(query, listItemId)

	err := row.Scan(&listItem.Title, &listItem.Id, &listItem.Notes, &listItem.UserId)

	if err != nil {
		return listItem, &models.NotFoundError{}
	}

	if listItem.UserId != userId {
		return models.ListItem{}, &models.AuthenticationError{}
	}

	return listItem, err
}
