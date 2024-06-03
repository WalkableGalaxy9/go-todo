package main

import (
	"log"
	"net/http"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

func main() {
	handler := http.HandlerFunc(gotodo.TodoServer)
	log.Fatal(http.ListenAndServe(":5000", handler))

}
