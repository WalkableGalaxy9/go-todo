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
	fmt.Fprintf(output, "\n-----MENU-----\nA. Add Todo\nD. Delete\n")
}

func ConvertInputToMenuOption(option rune) MenuOption {

	option = unicode.ToUpper(option)

	switch option {
	case 'A':
		return MenuAddTodo
	case 'D':
		return MenuDelete
	default:
		return MenuUnknown
	}
}
