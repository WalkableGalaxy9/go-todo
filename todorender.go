package gotodo

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*.gohtml"
	todoTemplates embed.FS
)

type TodoRender interface {
	RenderTodo(output io.Writer) error
	RenderCreate(output io.Writer) error
	CreateTodoFromPage(newTodo string) error
	DeleteTodoFromPage(deleteTodoIndex int) error
}

type TodoRenderHTML struct {
	TodoList TodoList
}

func (r *TodoRenderHTML) RenderTodo(output io.Writer) error {

	todoTemplate, err := template.New("Todo").Funcs(template.FuncMap{
		"statusString": StatusString}).ParseFS(todoTemplates, "templates/*.gohtml")

	if err != nil {
		return err
	}

	err = todoTemplate.ExecuteTemplate(output, "todo.gohtml", r.TodoList)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRenderHTML) RenderCreate(output io.Writer) error {

	todoTemplate, err := template.New("Todo").Funcs(template.FuncMap{
		"statusString": StatusString}).ParseFS(todoTemplates, "templates/*.gohtml")

	if err != nil {
		return err
	}

	err = todoTemplate.ExecuteTemplate(output, "new.gohtml", r.TodoList)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRenderHTML) CreateTodoFromPage(newTodo string) error {

	r.TodoList.CreateTodo(newTodo)
	return nil
}

func StatusString(status bool) string {
	if !status {
		return "incomplete"
	} else {
		return "complete"
	}

}

func (r *TodoRenderHTML) DeleteTodoFromPage(deleteTodoIndex int) error {

	return r.TodoList.DeleteTodo(deleteTodoIndex)
}
