package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/services"
	"github.com/dshoulders/goapi/utils"
)

func GetListItem(w http.ResponseWriter, r *http.Request) {
	itemId, err := utils.GetRequestParamAsInt("itemId", r)
	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if err != nil {
		response := models.CreateErrorResponse("List item Id is required")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusBadRequest)
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

func CreateListItem(w http.ResponseWriter, r *http.Request) {
	var listItem models.ListItem
	listId, err := utils.GetRequestParamAsInt("listId", r)

	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err = json.NewDecoder(r.Body).Decode(&listItem)
	if err != nil {
		response := models.CreateErrorResponse(err.Error())
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	listItem.UserId = userId

	listItem, err = services.SaveListItem(listId, listItem)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(listItem)
			utils.Respond(w, response, http.StatusOK)
			return
		}

	default:
		{
			response := models.CreateErrorResponse("Cannot save the list item")
			utils.Respond(w, response, http.StatusInternalServerError)
			return
		}
	}
}

func UpdateListItem(w http.ResponseWriter, r *http.Request) {
	var listItem models.ListItem
	listItemId, err := utils.GetRequestParamAsInt("itemId", r)

	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err = json.NewDecoder(r.Body).Decode(&listItem)
	if err != nil {
		response := models.CreateErrorResponse(err.Error())
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	listItem.UserId = userId
	listItem.Id = listItemId

	listItem, err = services.UpdateListItem(listItem)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(listItem)
			utils.Respond(w, response, http.StatusOK)
			return
		}

	case *models.NotFoundError:
		{
			response := models.CreateErrorResponse("List item was not found")
			utils.Respond(w, response, http.StatusInternalServerError)
			return
		}

	default:
		{
			response := models.CreateErrorResponse("Cannot save the list item")
			utils.Respond(w, response, http.StatusInternalServerError)
			return
		}
	}
}

func DeleteListItem(w http.ResponseWriter, r *http.Request) {

	listItemId, err := utils.GetRequestParamAsInt("itemId", r)

	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if err != nil {
		response := models.CreateErrorResponse("List item Id is required")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	ok, err = services.DeleteListItem(userId, listItemId)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(&models.Message{"List item was successfully deleted"})
			utils.Respond(w, response, http.StatusOK)
			return
		}

	case *models.NotFoundError:
		{
			response := models.CreateErrorResponse("List item was not found")
			utils.Respond(w, response, http.StatusInternalServerError)
			return
		}

	default:
		{
			response := models.CreateErrorResponse("Cannot delete the list item")
			utils.Respond(w, response, http.StatusInternalServerError)
			return
		}
	}
}
