package gotodo

import (
	"fmt"
	"math/rand/v2"
)

func OddRoutine() chan uint {
	channel := make(chan uint)
	go func() {
		oddNumbers := []uint{1, 3, 5, 7, 9}

		// Select a random index from the oddNumbers slice
		randomIndex := rand.IntN(len(oddNumbers))

		channel <- oddNumbers[randomIndex]
		close(channel)
	}()

	return channel
}

func EvenRoutine() chan uint {
	channel := make(chan uint)

	go func() {
		evenNumbers := []uint{2, 4, 6, 8, 10}

		// Select a random index from the oddNumbers slice
		randomIndex := rand.IntN(len(evenNumbers))

		channel <- evenNumbers[randomIndex]
		close(channel)
	}()

	return channel
}

func Dunno() {
	// 14. Write a program to simulate a race condition occurring when one goroutine updates a data variable with odd numbers,
	//while another updates the same data variable with even numbers. After each update ,
	//attempt to display the data contained in the data variable to screen. [Goroutines][Concurrency][Race Conditions]

	//15. Refactor the program created in exercise 14 to use channels, mutexes to synchronise all actions.
	//[Concurrency][Waitgroups][Workerpools][Mutexes]
	for i := 0; i < 10; i++ {

		//var number = uint(0)
		select {
		case number := <-OddRoutine():
			fmt.Printf("Outside odd: %d\n", number)
		case number := <-EvenRoutine():
			fmt.Printf("Outside even: %d\n", number)
		}

		//fmt.Printf("Outside: %d\n", <-channel)
	}
}
