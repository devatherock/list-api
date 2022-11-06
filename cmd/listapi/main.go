package main

import (
	"net/http"
	"os"

	"github.com/devatherock/list-api/cmd/listapi/apis"
	_ "github.com/devatherock/list-api/cmd/listapi/docs"
	util "github.com/devatherock/list-api/cmd/listapi/utilities"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
)

// @title List API
// @version 1.0
// @description APIs to manage lists.
// @termsOfService http://swagger.io/terms/

// @contact.url https://github.com/devatherock

// @license.name MIT

// @host https://list-api.onrender.com
// @BasePath /
func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/health", apis.GetHealth).Methods("GET")

	listRouter := router.PathPrefix("/users/{userId}/lists").Subrouter()
	listRouter.HandleFunc("", util.LogHandler(apis.GetLists)).Methods("GET")
	listRouter.HandleFunc("", util.LogHandler(apis.CreateList)).Methods("POST")
	listRouter.HandleFunc("/{listId}", util.LogHandler(apis.GetList)).Methods("GET")
	listRouter.HandleFunc("/{listId}", util.LogHandler(apis.DeleteList)).Methods("DELETE")
	listRouter.HandleFunc("/{listId}", util.LogHandler(apis.UpdateList)).Methods("PUT")

	// Read from PORT environment variable available on heroku
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("Http server listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
