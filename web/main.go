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
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/new", newHandler)
	err := http.ListenAndServe("localhost:8081", nil)

	log.Fatal(err)
}

func interactHandler(writer http.ResponseWriter, request *http.Request) {

	todoHTML.RenderTodo(writer)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {

	todo := request.FormValue("todo")

	todoHTML.CreateTodoFromPage(todo)
	http.Redirect(writer, request, "/", http.StatusFound)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {

	todoHTML.RenderCreate(writer)
}
