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

func RenderTodo(output io.Writer, todoList TodoList) error {

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

	err = todoTemplate.ExecuteTemplate(output, "todo.gohtml", todoList)

	if err != nil {
		return err
	}

	return nil
}
