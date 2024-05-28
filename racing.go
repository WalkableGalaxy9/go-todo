package gotodo

import (
	"fmt"
	"math/rand/v2"
)

func OddRoutine(number *uint) {

	oddNumbers := []uint{1, 3, 5, 7, 9}

	// Select a random index from the oddNumbers slice
	randomIndex := rand.IntN(len(oddNumbers))

	*number = oddNumbers[randomIndex]

	fmt.Printf("Odd: %d\n", *number)
}

func EvenRoutine(number *uint) {

	evenNumbers := []uint{2, 4, 6, 8, 10}

	// Select a random index from the oddNumbers slice
	randomIndex := rand.IntN(len(evenNumbers))

	*number = evenNumbers[randomIndex]

	fmt.Printf("Even: %d\n", *number)
}

func Dunno() {
	// 14. Write a program to simulate a race condition occurring when one goroutine updates a data variable with odd numbers,
	//while another updates the same data variable with even numbers. After each update ,
	//attempt to display the data contained in the data variable to screen. [Goroutines][Concurrency][Race Conditions]

	for i := 0; i < 10; i++ {

		var number = uint(0)
		go OddRoutine(&number)
		go EvenRoutine(&number)

		fmt.Printf("Outside: %d\n", number)
	}
}
