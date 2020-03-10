package services

import (
	"log"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/utils"
)

func GetListItem(listItemId int, userId int32) (models.ListItem, error) {

	var listItem models.ListItem

	query := "SELECT title, id, notes, app_user_id" +
		" FROM list_item" +
		" WHERE id = $1"

	dbConn := utils.GetDBConnection()
	defer dbConn.Close()

	rows, err := dbConn.Query(query, listItemId)
	if err != nil {
		log.Fatal(err)
		return listItem, err
	}

	for rows.Next() {
		err := rows.Scan(&listItem.Title, &listItem.Id, &listItem.Notes, &listItem.UserId)
		if err != nil {
			log.Fatal(err)
			return listItem, err
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return listItem, err
	}

	return listItem, err
}
