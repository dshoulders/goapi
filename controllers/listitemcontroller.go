package controllers

import (
	"net/http"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/services"
	"github.com/dshoulders/goapi/utils"
)

func ListItem(w http.ResponseWriter, r *http.Request) {
	itemId, err := utils.GetRequestParamAsInt("itemId", r)
	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if err != nil {
		response := models.CreateErrorResponse("List item Id is required")
		utils.Respond(w, response, http.StatusInternalServerError)
		return
	}

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusInternalServerError)
		return
	}

	listItem, err := services.GetListItem(itemId, userId)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(listItem)
			utils.Respond(w, response, http.StatusOK)
		}

	case *models.NotFoundError:
		{
			response := models.CreateErrorResponse("List item not found")
			utils.Respond(w, response, http.StatusNotFound)
			return
		}

	case *models.AuthenticationError:
		{
			response := models.CreateErrorResponse("Not authorised to access list item")
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
