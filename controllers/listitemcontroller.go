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
		response := models.CreateErrorResponse("need a list item Id")
		utils.Respond(w, response, http.StatusInternalServerError)
		return
	}

	if ok == false {
		response := models.CreateErrorResponse("user Id is not valid")
		utils.Respond(w, response, http.StatusInternalServerError)
		return
	}

	listItem, err := services.GetListItem(itemId, userId)

	if err != nil {
		response := models.CreateErrorResponse(err.Error())
		utils.Respond(w, response, http.StatusInternalServerError)
		return
	}

	response := models.CreateSuccessResponse(listItem)
	utils.Respond(w, response, http.StatusOK)
}
