package gotodo

import (
	"bytes"
	"testing"
)

func TestTopMenu(t *testing.T) {

	t.Run("Add todo", func(t *testing.T) {

		output := bytes.Buffer{}
		input := bytes.Buffer{}
		input.WriteString("A")

		got := GetMenuOption(&input, &output)
		gotMenu := output.String()

		AssertMenuOption(got, MenuAddTodo, t)
		AssertString(gotMenu, MenuString, t)
	})

	t.Run("Unknown", func(t *testing.T) {
		output := bytes.Buffer{}
		input := bytes.Buffer{}
		input.WriteString(".")

		got := GetMenuOption(&input, &output)
		gotMenu := output.String()

		AssertMenuOption(got, MenuUnknown, t)
		AssertString(gotMenu, MenuString, t)
	})
}

func TestConvertInputToMenuOption(t *testing.T) {

	var cases = []struct {
		Name   string
		Input  rune
		Option MenuOption
	}{
		{"Add", 'A', MenuAddTodo},
		{"Add lower", 'a', MenuAddTodo},
		{"None", ' ', MenuUnknown},
		{"Dot", '.', MenuUnknown},
		{"Delete", 'D', MenuDelete},
		{"Delete lower", 'd', MenuDelete},
		{"Toggle", 'T', MenuToggle},
		{"Toggle lower", 't', MenuToggle},
		{"Exit", 'X', MenuExit},
		{"Exit lower", 'x', MenuExit},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {

			got := ConvertInputToMenuOption(test.Input)

			AssertMenuOption(got, test.Option, t)
		})
	}

}

func TestDisplayMenu(t *testing.T) {

	output := bytes.Buffer{}

	DisplayMenu(&output)

	gotMenu := output.String()

	AssertString(gotMenu, MenuString, t)
}
