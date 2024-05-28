package gotodo

import "math/rand/v2"

func OddRoutine() uint {

	oddNumbers := []int{1, 3, 5, 7, 9}

	// Select a random index from the oddNumbers slice
	randomIndex := rand.IntN(len(oddNumbers))

	return uint(oddNumbers[randomIndex])
}

func EvenRoutine() uint {

	evenNumbers := []int{2, 4, 6, 8, 10}

	// Select a random index from the oddNumbers slice
	randomIndex := rand.IntN(len(evenNumbers))

	return uint(evenNumbers[randomIndex])
}
