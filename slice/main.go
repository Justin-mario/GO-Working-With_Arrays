package main

import (
	"fmt"
	"math/rand"
)

func main() {
	singleSlice()
	multidimensionalSlice()
}

func singleSlice() {
	var numbers = make([]int, 4)
	for i := 0; i < 4; i++ {
		numbers[i] = rand.Intn(100)
	}
	fmt.Println(numbers)
}

func multidimensionalSlice() {
	var numbers = make([][]int, 4)
	for i := 0; i < 4; i++ {
		numbers[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			numbers[i][j] = rand.Intn(400)
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(" ", +numbers[i][j])
		}
	}

}
