package gotodo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"unicode"
)

type MenuOption int

const (
	MenuUnknown = iota
	MenuAddTodo
	MenuDelete
	MenuToggle
	MenuExit
)

const (
	MenuString = "\n-----MENU-----\nA. Add Todo\nD. Delete\nT. Toggle\nX. Exit\n"
)

func GetMenuOption(input io.Reader, output io.Writer) MenuOption {

	DisplayMenu(output)

	reader := bufio.NewReader(input)
	option, _, err := reader.ReadRune()

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}

	return ConvertInputToMenuOption(option)
}

func DisplayMenu(output io.Writer) {
	fmt.Fprintf(output, MenuString)
}

func ConvertInputToMenuOption(option rune) MenuOption {

	option = unicode.ToUpper(option)

	switch option {
	case 'A':
		return MenuAddTodo
	case 'D':
		return MenuDelete
	case 'T':
		return MenuToggle
	case 'X':
		return MenuExit
	default:
		return MenuUnknown
	}
}
