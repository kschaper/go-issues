package main

import (
	"log"
	"net/http"

	"github.com/kschaper/go-issues/big-interfaces/db"
	"github.com/kschaper/go-issues/big-interfaces/handler"
)

func init() {
	var (
		dbClient   = &db.DatabaseClient{}
		userGetter = &db.UserService{DatabaseClient: dbClient}
	)
	http.HandleFunc("/users/", handler.GetUserHandler(userGetter))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
