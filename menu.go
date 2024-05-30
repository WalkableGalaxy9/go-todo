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
	MenuExit
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
	fmt.Fprintf(output, "\n-----MENU-----\nA. Add Todo\nD. Delete\nX. Exit\n")
}

func ConvertInputToMenuOption(option rune) MenuOption {

	option = unicode.ToUpper(option)

	switch option {
	case 'A':
		return MenuAddTodo
	case 'D':
		return MenuDelete

	case 'X':
		return MenuExit
	default:
		return MenuUnknown
	}
}
