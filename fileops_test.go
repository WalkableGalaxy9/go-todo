package gotodo

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"testing/fstest"
)

type TestFileSystemOps struct {
	// Keep a view on the input so we can test it
	Input bytes.Buffer
}

func (t *TestFileSystemOps) OpenFile(filename string) (io.Writer, error) {
	t.Input = bytes.Buffer{}
	return &t.Input, nil
}

func TestWriteJSONToFile(t *testing.T) {

	fileSystem := TestFileSystemOps{}

	todolist := TodoList{
		{"Do laundry", true},
		{"Go shopping", false},
		{"learn go", false},
	}

	want := "[{\"Title\":\"Do laundry\",\"Complete\":true},{\"Title\":\"Go shopping\",\"Complete\":false},{\"Title\":\"learn go\",\"Complete\":false}]\n"
	filePath := "output/todo.txt"
	err := WriteJSONToFile(&fileSystem, filePath, todolist)

	if err != nil {
		t.Errorf("Unexpected error %d", err)
	}

	got := fileSystem.Input.String()
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func TestReadJSONFromAFile(t *testing.T) {

	filename := "input/todo.txt"

	filesystem := fstest.MapFS{
		filename: {
			Data: []byte("[{\"Title\":\"Do laundry\",\"Complete\":true},{\"Title\":\"Go shopping\",\"Complete\":false},{\"Title\":\"learn go\",\"Complete\":false}]\n"),
		},
	}

	todolist, err := ReadJSONFromAFile(filesystem, filename)

	want := TodoList{
		{"Do laundry", true},
		{"Go shopping", false},
		{"learn go", false},
	}

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, todolist) {
		t.Errorf("Want %v got %v", want, todolist)
	}
}
