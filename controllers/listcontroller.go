package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/services"
	"github.com/dshoulders/goapi/utils"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	listId, err := utils.GetRequestParamAsInt("listId", r)
	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if err != nil {
		response := models.CreateErrorResponse("List Id is required")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	list, err := services.GetList(listId, userId)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(list)
			utils.Respond(w, response, http.StatusOK)
			return
		}

	case *models.NotFoundError:
		{
			response := models.CreateErrorResponse("List not found")
			utils.Respond(w, response, http.StatusNotFound)
			return
		}

	case *models.AuthenticationError:
		{
			response := models.CreateErrorResponse("Not authorised to access list")
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

func GetLists(w http.ResponseWriter, r *http.Request) {
	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusInternalServerError)
		return
	}

	lists, err := services.GetLists(userId)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(lists)
			utils.Respond(w, response, http.StatusOK)
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

func CreateList(w http.ResponseWriter, r *http.Request) {
	var list models.List

	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		response := models.CreateErrorResponse(err.Error())
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	list, err = services.SaveList(list.Title, userId)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(list)
			utils.Respond(w, response, http.StatusOK)
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

func GetListItems(w http.ResponseWriter, r *http.Request) {
	listId, err := utils.GetRequestParamAsInt("listId", r)
	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if err != nil {
		response := models.CreateErrorResponse("List Id is required")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusBadRequest)
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

func DeleteList(w http.ResponseWriter, r *http.Request) {
	listId, err := utils.GetRequestParamAsInt("listId", r)
	userValue := r.Context().Value("user")
	userId, ok := userValue.(int32)

	if err != nil {
		response := models.CreateErrorResponse("List Id is required")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	if ok == false {
		response := models.CreateErrorResponse("User Id is not valid")
		utils.Respond(w, response, http.StatusBadRequest)
		return
	}

	ok, err = services.DeleteList(userId, listId)

	switch err.(type) {
	case nil:
		{
			response := models.CreateSuccessResponse(&models.Message{"List was successfully deleted"})
			utils.Respond(w, response, http.StatusOK)
			return
		}

	case *models.NotFoundError:
		{
			response := models.CreateErrorResponse("List not found. List Id: " + strconv.Itoa(listId))
			utils.Respond(w, response, http.StatusNotFound)
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
