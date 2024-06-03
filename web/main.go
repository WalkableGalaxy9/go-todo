package main

import (
	"log"
	"net/http"
	"strconv"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

var todoHTML gotodo.TodoRenderHTML = gotodo.TodoRenderHTML{}

func main() {

	//todoHTML := gotodo.TodoRenderHTML{}
	todoHTML.TodoList = gotodo.TodoList{
		{
			Title:    "Do laundry",
			Complete: false,
		},
	}

	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/delete/{id}", deleteHandler)
	http.HandleFunc("/", interactHandler)
	err := http.ListenAndServe("localhost:8081", nil)

	log.Fatal(err)
}

func interactHandler(writer http.ResponseWriter, request *http.Request) {
	HandleRender(&todoHTML, writer)
}

func HandleRender(renderer gotodo.TodoRender, writer http.ResponseWriter) {
	renderer.RenderTodo(writer)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {

	HandleCreate(&todoHTML, request, writer)
}

func HandleCreate(renderer gotodo.TodoRender, request *http.Request, writer http.ResponseWriter) {
	todo := request.FormValue("todo")

	renderer.CreateTodoFromPage(todo)
	http.Redirect(writer, request, "/", http.StatusFound)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {

	HandleCreateRender(&todoHTML, writer)
}

func HandleCreateRender(renderer gotodo.TodoRender, writer http.ResponseWriter) {
	renderer.RenderCreate(writer)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request) {

	// obtain the index to delete from the url
	// the index we get from the page is 0-based whereas the delete function expects a 1-based
	HandleDelete(&todoHTML, request, writer)
}

func HandleDelete(renderer gotodo.TodoRender, request *http.Request, writer http.ResponseWriter) {
	indexStr := request.URL.Path[len("/delete/"):]
	index, _ := strconv.Atoi(indexStr)

	index++

	renderer.DeleteTodoFromPage(index)
	http.Redirect(writer, request, "/", http.StatusFound)
}
