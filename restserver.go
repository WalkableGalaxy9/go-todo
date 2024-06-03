package gotodo

import (
	"fmt"
	"net/http"
)

var todoList TodoList = TodoList{}

func TodoServer(writer http.ResponseWriter, request *http.Request) {

	PrintTodoJSON(writer, todoList)

	fmt.Fprint(writer, "[{\"Title\":\"Do laundry\",\"Complete\":true},{\"Title\":\"Go shopping\",\"Complete\":false},{\"Title\":\"learn go\",\"Complete\":false}]\n")
}
