package gotodo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type FileOps interface {
	// Brought this into an interface so we can mock it for testing
	OpenFile(filename string) (io.Writer, error)
}

type FileSystemOps struct {
}

func (f *FileSystemOps) OpenFile(filename string) (io.Writer, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return file, nil
}

func WriteJSONToFile(fileOperations FileOps, filename string, todoList TodoList) error {

	file, err := fileOperations.OpenFile(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return err
	}

	PrintTodoJSON(file, todoList)

	return nil
}

func ReadJSONFromAFile(fileSystem fs.FS, filename string) (todoList TodoList, err error) {
	todoFile, err := fileSystem.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return nil, err
	}
	defer todoFile.Close()

	fileData, err := io.ReadAll(todoFile)

	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil, err
	}

	err = json.Unmarshal(fileData, &todoList)

	if err != nil {
		fmt.Printf("Error parsing JSON: %v", err)
		return nil, err
	}

	return todoList, nil
}

func PrintTodoJSON(output io.Writer, todoList TodoList) {

	json.NewEncoder(output).Encode(todoList)
}
