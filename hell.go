package main

import (
	"fmt"
)

func main() {
	numbers := []int{-6, -5, -67, -12, -54}
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	fmt.Println(max)
}
