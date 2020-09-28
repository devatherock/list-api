package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/devatherock/list-api/cmd/listapi/models"
	"github.com/devatherock/list-api/cmd/listapi/services"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// GetLists godoc
// @Summary Gets all lists for a user
// @Description Gets all lists for a user
// @Produce json
// @Param userId path string true "The user id of the user"
// @Success 200 {array} models.List
// @Failure 404 "The user does not have any lists"
// @Router /user/{userId}/lists [get]
func GetLists(writer http.ResponseWriter, request *http.Request) {
	userId := readUserId(request)

	lists := services.GetLists(userId)
	if lists != nil {
		writeResponse(writer, lists)
	} else {
		writeStatus(writer, 404)
	}
}

// CreateList godoc
// @Summary Creates a list
// @Description Creates a list
// @Accept json
// @Produce json
// @Param userId path string true "The user id of the user"
// @Param list body models.List true "The value of the list"
// @Success 201 {object} models.List "The created list"
// @Router /user/{userId}/lists [post]
func CreateList(writer http.ResponseWriter, request *http.Request) {
	// Read the list
	list, error := readList(writer, request)

	if error == nil {
		userId := readUserId(request)
		writeResponseWithStatus(writer, services.CreateList(userId, list), 201)
	}
}

// GetList godoc
// @Summary Gets a specific list of a user
// @Description Gets a specific list of a user
// @Produce json
// @Param userId path string true "The user id of the user"
// @Param listId path string true "The id of the list to get"
// @Success 200 {object} models.List
// @Failure 404 "list with the specified id does not exist"
// @Router /user/{userId}/lists/{listId} [get]
func GetList(writer http.ResponseWriter, request *http.Request) {
	userId, listId := readUserIdAndListId(request)

	list := services.GetList(userId, listId)
	if list != nil {
		writeResponse(writer, list)
	} else {
		writeStatus(writer, 404)
	}
}

// UpdateList godoc
// @Summary Updates a specific list of a user
// @Description Updates a specific list of a user
// @Accept json
// @Param userId path string true "The user id of the user"
// @Param listId path string true "The id of the list to update"
// @Param list body models.List true "The new value of the list"
// @Success 202 "updated successfully"
// @Failure 404 "list with the specified id does not exist"
// @Router /user/{userId}/lists/{listId} [put]
func UpdateList(writer http.ResponseWriter, request *http.Request) {
	// Read the list
	list, error := readList(writer, request)

	if error == nil {
		userId, listId := readUserIdAndListId(request)

		updateResult := services.UpdateList(userId, listId, list)
		if updateResult {
			writeStatus(writer, 202)
		} else {
			writeStatus(writer, 404)
		}
	}
}

// DeleteList godoc
// @Summary Deletes a specific list of a user
// @Description Deletes a specific list of a user
// @Param userId path string true "The user id of the user"
// @Param listId path string true "The id of the list to update"
// @Success 204 "deleted successfully"
// @Failure 404 "list with the specified id does not exist"
// @Router /user/{userId}/lists/{listId} [delete]
func DeleteList(writer http.ResponseWriter, request *http.Request) {
	userId, listId := readUserIdAndListId(request)

	deleteResult := services.DeleteList(userId, listId)
	if deleteResult {
		writeStatus(writer, 204)
	} else {
		writeStatus(writer, 404)
	}
}

// Reads userId path variable
func readUserId(request *http.Request) (userId string) {
	pathVariables := mux.Vars(request)
	userId = pathVariables["userId"]

	return
}

// Reads userId and listId path variable
func readUserIdAndListId(request *http.Request) (userId string, listId string) {
	pathVariables := mux.Vars(request)
	userId = pathVariables["userId"]
	listId = pathVariables["listId"]

	return
}

// Reads a list from the request body
func readList(writer http.ResponseWriter, request *http.Request) (*models.List, error) {
	// Read request
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writeErrorResponse(writer, err, 400)
		return nil, err
	}

	// Parse request
	list := models.List{}
	err = json.Unmarshal(requestBody, &list)
	if err != nil {
		writeErrorResponse(writer, err, 400)
		return nil, err
	}

	return &list, nil
}

// Writes an error HTTP status code
func writeErrorResponse(writer http.ResponseWriter, err error, status int) {
	log.Error("error: ", err)
	writeResponseWithStatus(writer, map[string]interface{}{
		"error": err.Error(),
	}, status)
}

// Writes the response with 200 status code
func writeResponse(writer http.ResponseWriter, response interface{}) {
	writeResponseWithStatus(writer, response, 200)
}

// Write the response with the specified HTTP status code
func writeResponseWithStatus(writer http.ResponseWriter, response interface{}, status int) {
	writeStatus(writer, status)
	writeResponseBody(writer, response)
}

// Write the response with the specified HTTP status code
func writeStatus(writer http.ResponseWriter, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
}

// Writes the response body
func writeResponseBody(writer http.ResponseWriter, response interface{}) {
	responseBody, error := json.Marshal(&response)
	if error != nil {
		log.Error("error: ", error)
		return
	} else {
		writer.Write(responseBody)
	}
}
