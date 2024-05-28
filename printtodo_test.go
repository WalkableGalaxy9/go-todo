package gotodo

import (
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestPrintTodo(t *testing.T) {

	t.Run("Testing Incomplete Item", func(t *testing.T) {
		todo := []TodoItem{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer, todo)

		want := "1. Do laundry - incomplete\n2. Go shopping - incomplete\n3. learn go - incomplete\n"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
	t.Run("Testing complete Item", func(t *testing.T) {
		todo := []TodoItem{
			{"Do laundry", true},
			{"Go shopping", true},
			{"learn go", true},
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer, todo)

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

	got := CreateList("Do laundry", "Go shopping", "learn go")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %v got %v", want, got)
	}
}

func TestPrintJSON(t *testing.T) {
	todo := []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	buffer := bytes.Buffer{}
	PrintTodoJSON(&buffer, todo)

	//response := buffer.String()
	var got []TodoItem
	err := json.NewDecoder(&buffer).Decode(&got)

	if err != nil {
		t.Errorf("Unable to convert response to JSON: %s, %s", buffer.String(), err)
	}
}

func TestWriteJSONToFile(t *testing.T) {
	items := []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	filePath := "output/todo.txt"
	err := WriteJSONToFile(filePath, items)

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
