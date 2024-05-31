package gotodo

import (
	"reflect"
	"testing"
)

func AssertTodoList(want TodoList, got TodoList, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v got %v", want, got)
	}
}

func AssertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Got an error %d", err)
	}
}

func AssertError(err error, t *testing.T) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected an error but didn't get one")
	}
}

func AssertString(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("Want %s got %s", want, got)
	}
}

func AssertMenuOption(got, want MenuOption, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v wanted %v", got, want)
	}
}
