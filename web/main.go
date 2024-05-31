package main

import (
	"log"
	"net/http"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

var todoHTML gotodo.TodoRenderHTML

func main() {
	todoHTML.TodoList = gotodo.TodoList{
		{
			Title:    "Do laundry",
			Complete: false,
		},
	}

	http.HandleFunc("/", interactHandler)
	err := http.ListenAndServe("localhost:8081", nil)

	log.Fatal(err)
}

func interactHandler(writer http.ResponseWriter, request *http.Request) {

	todoHTML.RenderTodo(writer)
}
