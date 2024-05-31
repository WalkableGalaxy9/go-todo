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
	RenderTodo(output io.Writer, todoList TodoList) error
}

type TodoRenderHTML struct {
	TodoList TodoList
}

func (r *TodoRenderHTML) RenderTodo(output io.Writer) error {

	todoTemplate, err := template.New("Todo").Funcs(template.FuncMap{
		"statusString": func(status bool) string {
			if !status {
				return "incomplete"
			} else {
				return "complete"
			}
		},
	}).ParseFS(todoTemplates, "templates/*.gohtml")

	if err != nil {
		return err
	}

	err = todoTemplate.ExecuteTemplate(output, "todo.gohtml", r.TodoList)

	if err != nil {
		return err
	}

	return nil
}
