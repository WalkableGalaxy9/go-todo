package gotodo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodo(t *testing.T) {
	t.Run("Returns the list of Todos", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todo", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := "[{\"Title\":\"Do laundry\",\"Complete\":true},{\"Title\":\"Go shopping\",\"Complete\":false},{\"Title\":\"learn go\",\"Complete\":false}]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
