package services

import (
	"github.com/devatherock/list-api/cmd/listapi/models"
	"github.com/google/uuid"
)

var allLists = make(map[string]map[string]*models.List)

// Retrieves all lists of an user
func GetLists(userId string) (userLists []*models.List) {
	listsByUser := allLists[userId]
	if listsByUser != nil {
		for _, list := range listsByUser {
			userLists = append(userLists, list)
		}
	}

	return
}

// Creates a list
func CreateList(userId string, list *models.List) *models.List {
	list.Id = uuid.New().String()

	listsByUser := allLists[userId]
	if listsByUser == nil {
		listsByUser = make(map[string]*models.List)
		allLists[userId] = listsByUser
	}

	listsByUser[list.Id] = list
	return list
}

// Retrieves a list
func GetList(userId string, listId string) (list *models.List) {
	listsByUser := allLists[userId]
	if listsByUser != nil {
		list = listsByUser[listId]
	}

	return
}

// Updates a list
func UpdateList(userId string, listId string, list *models.List) (updated bool) {
	listsByUser := allLists[userId]

	if listsByUser != nil {
		listToUpdate := listsByUser[listId]
		if listToUpdate != nil {
			list.Id = listId
			listsByUser[listId] = list
			updated = true
		}
	}

	return
}

// Deletes a list
func DeleteList(userId string, listId string) (deleted bool) {
	listsByUser := allLists[userId]

	if listsByUser != nil {
		listToDelete := listsByUser[listId]
		if listToDelete != nil {
			delete(listsByUser, listId)
			deleted = true
		}
	}

	return
}
