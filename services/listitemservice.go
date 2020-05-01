package services

import (
	"log"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/utils"
)

func getListOwner(listId int) (int, error) {
	query := `
		SELECT app_user_id
		FROM list
		WHERE list_id = $1
	`

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

	query := `
		SELECT li.title, li.item_id, li.notes, au.id
		FROM list_item li
		INNER JOIN app_user au
		   on li.app_user_id = au.app_user_id
		WHERE list_id = $1
	`

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
			WHERE list_id = $3 AND app_user_id = $4
		)
		RETURNING item_id
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

func UpdateListItem(listItem models.ListItem) (models.ListItem, error) {

	query := `
		UPDATE list_item
		SET title = $2, notes = $3		
		WHERE item_id = $1 AND app_user_id = $4
		`

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	res, err := dbConn.Exec(query, listItem.Id, listItem.Title, listItem.Notes, listItem.UserId)

	if err != nil {
		return listItem, err
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return listItem, &models.NotFoundError{}
	}

	return listItem, nil
}

func DeleteListItem(userId int32, listItemId int) (bool, error) {

	query := `
		DELETE from list_item	
		WHERE item_id = $1 AND app_user_id = $2
		`

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	res, err := dbConn.Exec(query, listItemId, userId)

	if err != nil {
		return false, err
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return false, &models.NotFoundError{}
	}

	return true, nil
}

func GetListItem(listItemId int, userId int32) (models.ListItem, error) {

	var listItem models.ListItem

	query := `
		SELECT title, item_id, notes, app_user_id
		FROM list_item
		WHERE item_id = $1
	`

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
