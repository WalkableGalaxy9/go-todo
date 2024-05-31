package gotodo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type FileOps interface {
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

func printTodoJSON(writer io.Writer, items TodoList) {

	json.NewEncoder(writer).Encode(items)
}

func WriteJSONToFile(fs FileOps, filename string, items TodoList) error {

	file, err := fs.OpenFile(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return err
	}

	printTodoJSON(file, items)

	return nil
}

func ReadJSONFromAFile(filesystem fs.FS, filename string) (items []TodoItem, err error) {
	todofile, err := filesystem.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return nil, err
	}
	defer todofile.Close()

	data, err := io.ReadAll(todofile)

	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &items)

	if err != nil {
		fmt.Printf("Error parsing JSON: %v", err)
		return nil, err
	}

	return items, nil
}
