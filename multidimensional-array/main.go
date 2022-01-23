package main

import (
	"fmt"
	"math/rand"
)

func main() {
	matrix()
}

func matrix() {
	var numbers [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			numbers[i][j] = rand.Intn(200)
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(numbers[i][j])
		}
	}

}
