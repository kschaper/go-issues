package main

import (
	"log"
	"net/http"

	"github.com/kschaper/go-issues/big-interfaces"

	"github.com/kschaper/go-issues/big-interfaces/db"
	"github.com/kschaper/go-issues/big-interfaces/handler"
)

func init() {
	env := whatever.Env{DB: &db.DatabaseClient{}}
	http.HandleFunc("/users/", handler.GetUserHandler(env))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
