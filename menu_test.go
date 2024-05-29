package gotodo

import (
	"bytes"
	"testing"
)

func TestTopMenu(t *testing.T) {

	t.Run("Add todo", func(t *testing.T) {

		output := bytes.Buffer{}
		input := bytes.Buffer{}
		wantmenu := "\n-----MENU-----\nA. Add Todo\nD. Delete\n"
		input.WriteString("A")

		got := GetMenuOption(&input, &output)

		if got != MenuAddTodo {
			t.Errorf("Got %v wanted %v", got, MenuAddTodo)
		}

		gotmenu := output.String()

		if gotmenu != wantmenu {
			t.Errorf("Got %v wanted %v", gotmenu, wantmenu)
		}
	})

	t.Run("Unknown", func(t *testing.T) {
		output := bytes.Buffer{}
		input := bytes.Buffer{}
		wantmenu := "\n-----MENU-----\nA. Add Todo\nD. Delete\n"
		input.WriteString(".")

		got := GetMenuOption(&input, &output)

		if got != MenuUnknown {
			t.Errorf("Got %v wanted %v", got, MenuUnknown)
		}

		gotmenu := output.String()

		if gotmenu != wantmenu {
			t.Errorf("Got %v wanted %v", gotmenu, wantmenu)
		}
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
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {

			got := ConvertInputToMenuOption(test.Input)

			if got != test.Option {
				t.Errorf("Got %v wanted %v for input %c", got, test.Option, test.Input)
			}
		})
	}

}

func TestDisplayMenu(t *testing.T) {

	output := bytes.Buffer{}
	wantmenu := "\n-----MENU-----\nA. Add Todo\nD. Delete\n"

	DisplayMenu(&output)

	gotmenu := output.String()

	if gotmenu != wantmenu {
		t.Errorf("Got %v wanted %v", gotmenu, wantmenu)
	}

}
