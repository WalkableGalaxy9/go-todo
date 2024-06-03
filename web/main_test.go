package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyTodoRenderHTML struct {
}

var RenderCount = 0

func (r *SpyTodoRenderHTML) RenderTodo(output io.Writer) error {

	RenderCount++
	return nil
}

var RenderCreateCount = 0

func (r *SpyTodoRenderHTML) RenderCreate(output io.Writer) error {

	RenderCreateCount++
	return nil
}

var CreateTodoFromPageCount = 0

func (r *SpyTodoRenderHTML) CreateTodoFromPage(newTodo string) error {

	CreateTodoFromPageCount++
	return nil
}

var DeleteTodoFromPageCount = 0

func (r *SpyTodoRenderHTML) DeleteTodoFromPage(deleteTodoIndex int) error {

	DeleteTodoFromPageCount++
	return nil
}

var ToggleTodoFromPageCount = 0

func (r *SpyTodoRenderHTML) ToggleTodoFromPage(toggleTodoIndex int) error {

	ToggleTodoFromPageCount++
	return nil
}

func resetCounts() {
	DeleteTodoFromPageCount = 0
	CreateTodoFromPageCount = 0
	RenderCount = 0
	RenderCreateCount = 0
	ToggleTodoFromPageCount = 0
}

func TestDeleteHandler(t *testing.T) {

	resetCounts()
	spyRenderer := SpyTodoRenderHTML{}
	// Create a new HTTP request
	request, err := http.NewRequest("POST", "/delete/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	responseWriter := httptest.NewRecorder()

	HandleDelete(&spyRenderer, request, responseWriter)

	assertCalls(t, 1, 0, 0, 0, 0)
}

func TestCreateHandler(t *testing.T) {

	resetCounts()
	spyRenderer := SpyTodoRenderHTML{}

	// Create a new HTTP request
	request, err := http.NewRequest("POST", "/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	responseWriter := httptest.NewRecorder()

	HandleCreate(&spyRenderer, request, responseWriter)

	assertCalls(t, 0, 1, 0, 0, 0)
}

func TestNewHandler(t *testing.T) {

	resetCounts()
	spyRenderer := SpyTodoRenderHTML{}

	// Create a new HTTP request
	_, err := http.NewRequest("POST", "/new", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	responseWriter := httptest.NewRecorder()

	HandleCreateRender(&spyRenderer, responseWriter)

	assertCalls(t, 0, 0, 0, 1, 0)

}

func TestToggleHandler(t *testing.T) {

	resetCounts()
	spyRenderer := SpyTodoRenderHTML{}

	// Create a new HTTP request
	request, err := http.NewRequest("POST", "/toggle/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	responseWriter := httptest.NewRecorder()

	HandleToggle(&spyRenderer, request, responseWriter)

	assertCalls(t, 0, 0, 0, 0, 1)

}

func assertCalls(t *testing.T, wantDelete, wantCreate, wantRender, wantRenderCreate, wantToggle int) {
	t.Helper()

	assertCall(t, wantDelete, DeleteTodoFromPageCount, "Delete Todo")
	assertCall(t, wantCreate, CreateTodoFromPageCount, "Create Todo")
	assertCall(t, wantRender, RenderCount, "Render")
	assertCall(t, wantRenderCreate, RenderCreateCount, "Render Create")
	assertCall(t, wantToggle, ToggleTodoFromPageCount, "Toggle Todo")
}

func assertCall(t *testing.T, want, got int, name string) {
	t.Helper()
	if want != got {
		t.Errorf("Expected %d calls to %s, got %d calls", want, name, got)
	}
}
