package gotodo

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"testing/fstest"
)

const (
	testFileName = "output/todo.txt"
)

type TestFileSystemOps struct {
	// Keep a view on the input so we can test it
	Input bytes.Buffer
}

func (t *TestFileSystemOps) OpenFile(filename string) (io.Writer, error) {

	if filename == testFileName {
		t.Input = bytes.Buffer{}
		return &t.Input, nil
	} else {
		return nil, fmt.Errorf("Error reading file")
	}
}

func TestWriteJSONToFile(t *testing.T) {

	t.Run("Green Path", func(t *testing.T) {
		fileSystem := TestFileSystemOps{}

		todolist := TodoList{
			{"Do laundry", true},
			{"Go shopping", false},
			{"learn go", false},
		}

		want := "[{\"Title\":\"Do laundry\",\"Complete\":true},{\"Title\":\"Go shopping\",\"Complete\":false},{\"Title\":\"learn go\",\"Complete\":false}]\n"

		err := WriteJSONToFile(&fileSystem, testFileName, todolist)

		if err != nil {
			t.Errorf("Unexpected error %d", err)
		}

		got := fileSystem.Input.String()
		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("File doesn't exist", func(t *testing.T) {
		fileSystem := TestFileSystemOps{}

		todolist := TodoList{
			{"Do laundry", true},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := WriteJSONToFile(&fileSystem, "This file doesnt exist", todolist)

		AssertError(err, t)
	})

}

func TestReadJSONFromAFile(t *testing.T) {

	t.Run("Green Path", func(t *testing.T) {

		fileSystem := fstest.MapFS{
			testFileName: {
				Data: []byte("[{\"Title\":\"Do laundry\",\"Complete\":true},{\"Title\":\"Go shopping\",\"Complete\":false},{\"Title\":\"learn go\",\"Complete\":false}]\n"),
			},
		}

		todoList, err := ReadJSONFromAFile(fileSystem, testFileName)

		want := TodoList{
			{"Do laundry", true},
			{"Go shopping", false},
			{"learn go", false},
		}

		AssertNoError(err, t)
		AssertTodoList(want, todoList, t)
	})

	t.Run("Bad filename", func(t *testing.T) {

		fileSystem := fstest.MapFS{
			testFileName: {
				Data: []byte("[{\"Title\":\"Do laundry\",\"Complete\":true},{\"Title\":\"Go shopping\",\"Complete\":false},{\"Title\":\"learn go\",\"Complete\":false}]\n"),
			},
		}

		_, err := ReadJSONFromAFile(fileSystem, "This file doesnt exist.txt")

		AssertError(err, t)
	})

	t.Run("Bad data", func(t *testing.T) {

		filesystem := fstest.MapFS{
			testFileName: {
				Data: []byte("This data is rubbish"),
			},
		}

		_, err := ReadJSONFromAFile(filesystem, testFileName)

		AssertError(err, t)
	})
}
