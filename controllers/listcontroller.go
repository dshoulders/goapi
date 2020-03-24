package controllers

import (
	"net/http"
	"strconv"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/services"
	"github.com/dshoulders/goapi/utils"
)

func ListItems(w http.ResponseWriter, r *http.Request) {
	listId, err := utils.GetRequestParamAsInt("listId", r)
	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if err != nil {
		response := models.CreateErrorResponse("List Id is required")
		utils.Respond(w, response, http.StatusInternalServerError)
		return
	}

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusInternalServerError)
		return
	}

	listItems, err := services.GetListItems(listId, userId)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(listItems)
			utils.Respond(w, response, http.StatusOK)
			return
		}

	case *models.NotFoundError:
		{
			response := models.CreateErrorResponse("List not found. List Id: " + strconv.Itoa(listId))
			utils.Respond(w, response, http.StatusNotFound)
			return
		}

	case *models.AuthenticationError:
		{
			response := models.CreateErrorResponse("Not authorised to access list. List Id: " + strconv.Itoa(listId))
			utils.Respond(w, response, http.StatusUnauthorized)
			return
		}

	default:
		{
			response := models.CreateErrorResponse(err.Error())
			utils.Respond(w, response, http.StatusInternalServerError)
			return
		}
	}
}
