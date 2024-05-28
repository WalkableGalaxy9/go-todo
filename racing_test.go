package gotodo

import "testing"

func TestOddRoutine(t *testing.T) {

	// Call the Odd Routine 10 times and check all the results are odd between 1 and 10

	for i := 0; i < 100; i++ {
		got := OddRoutine()

		if got < 1 || got > 10 || got%2 == 0 {
			t.Errorf("Wanted an odd number got %d", got)
		} else {
			t.Logf("Got %d", got)
		}
	}

}

func TestEvenRoutine(t *testing.T) {

	// Call the Odd Routine 10 times and check all the results are odd between 1 and 10

	for i := 0; i < 100; i++ {
		got := EvenRoutine()

		if got < 1 || got > 10 || got%2 != 0 {
			t.Errorf("Wanted an even number got %d", got)
		} else {
			t.Logf("Got %d", got)
		}
	}

}
