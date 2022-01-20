package main

import (
	"fmt"
)

var names [10]int

func main() {
	addingNumbersToArray()
	addingNumbersToArrayUsingRange()
	copingArrayElementsUsingRange()
	copingArrayElements()
	sortingArrayElements()
}

func addingNumbersToArray() {
	for i := 0; i < len(names); i++ {
		names[i] = i
	}
	fmt.Printf("numbers in arrays are %v\n", names)
	fmt.Printf("length of array names is  %v\n", len(names))
}

func addingNumbersToArrayUsingRange() {
	for i := range names {
		names[i] = i
	}
	fmt.Printf("numbers in arrays are %v\n", names)

}

func copingArrayElementsUsingRange() {
	states := [5]string{"Lagos", "Kano", "Abuja", "Enugu", "Rivers"}
	var copiedStates = [5]string{}

	for i, sate := range states {
		copiedStates[i] = sate
	}
	fmt.Printf("copied states are %v", copiedStates)
}

func copingArrayElements() {
	states := [5]string{"Lagos", "Kano", "Abuja", "Enugu", "Rivers"}
	var copiedStates = [3]string{}

	for i := 0; i < 3; i++ {

		copiedStates[i] = states[i]

	}
	fmt.Printf("copied states are %v", copiedStates)
}

func sortingArrayElements() {
	numbers := [12]int{2, 0, 4, 9, 1, 5, 3, 6, 8, 7, -12, -1}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			temp := 0
			if numbers[j] < numbers[i] {
				temp = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp

			}
		}

	}
	fmt.Printf("sorted array %v", numbers)
}
