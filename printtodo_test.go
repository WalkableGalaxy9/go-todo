package gotodo

import (
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestPrintTodo(t *testing.T) {

	t.Run("Testing Incomplete Item", func(t *testing.T) {
		TodoList = []TodoItem{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer)

		want := "1. Do laundry - incomplete\n2. Go shopping - incomplete\n3. learn go - incomplete\n"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
	t.Run("Testing complete Item", func(t *testing.T) {
		TodoList = []TodoItem{
			{"Do laundry", true},
			{"Go shopping", true},
			{"learn go", true},
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer)

		want := "1. Do laundry - complete\n2. Go shopping - complete\n3. learn go - complete\n"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
}

func TestCreateList(t *testing.T) {

	want := []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	CreateList("Do laundry", "Go shopping", "learn go")

	if !reflect.DeepEqual(TodoList, want) {
		t.Errorf("Want %v got %v", want, TodoList)
	}
}

func TestPrintJSON(t *testing.T) {
	TodoList = []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	buffer := bytes.Buffer{}
	PrintTodoJSON(&buffer)

	//response := buffer.String()
	var got []TodoItem
	err := json.NewDecoder(&buffer).Decode(&got)

	if err != nil {
		t.Errorf("Unable to convert response to JSON: %s, %s", buffer.String(), err)
	}
}

func TestWriteJSONToFile(t *testing.T) {
	TodoList = []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	filePath := "output/todo.txt"
	err := WriteJSONToFile(filePath)

	if err != nil {
		t.Errorf("Unexpected error %d", err)
	}
	// Check if the file exists
	if _, err := os.Stat(filePath); err == nil {
		t.Log("File exists.")
	} else if os.IsNotExist(err) {
		t.Error("File does not exist.")
	} else {
		t.Error("Error checking file:", err)
	}
}

func TestReadJSONFromAFile(t *testing.T) {
	fs := fstest.MapFS{
		"input/todo.txt": {Data: []byte(`[{"Title":"Do laundry","Complete":true},{"Title":"Go shopping","Complete":false},{"Title":"learn go","Complete":false}]`)},
	}

	buffer := bytes.Buffer{}
	err := ReadJSONFromAFile(fs, "input/todo.txt", &buffer)

	want := "1. Do laundry - complete\n2. Go shopping - incomplete\n3. learn go - incomplete\n"

	if err != nil {
		t.Fatal(err)
	}

	if want != buffer.String() {
		t.Errorf("Want %s got %s", want, buffer.String())
	}
}

func TestCreateTodo(t *testing.T) {

	want := []TodoItem{
		{
			"Do laundry",
			false,
		},
		{
			"Learn Go",
			true,
		},
	}

	items := []TodoItem{
		{
			"Do laundry",
			false,
		},
	}

	got := append(items, CreateTodo("Learn Go", true))

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v got %v", want, got)
	}

}

func TestAddItem(t *testing.T) {

	output := bytes.Buffer{}
	input := bytes.Buffer{}
	input.WriteString("New todo")

	AddTodoInput(&input, &output)

	want := []TodoItem{{Title: "New todo", Complete: false}}

	if !reflect.DeepEqual(TodoList, want) {
		t.Errorf("Got %v want %v", TodoList, want)
	}

	if output.String() != "Title:\n" {
		t.Errorf("Got %v want %v", output.String(), "Title:\n")
	}
}

func TestDeleteItem(t *testing.T) {

	TodoList = []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	output := bytes.Buffer{}
	input := bytes.Buffer{}
	input.WriteString("1\n")

	DeleteTodoInput(&input, &output)

	want := []TodoItem{
		{"Go shopping", false},
		{"learn go", false},
	}

	if !reflect.DeepEqual(TodoList, want) {
		t.Errorf("Got %v want %v", TodoList, want)
	}

	if output.String() != "Number:\n" {
		t.Errorf("Got %v want %v", output.String(), "Number:\n")
	}
}
