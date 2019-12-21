package main

import (
	"fmt"
)

func handlePanic() {
	text := recover()
	fmt.Println(text)
}
func Findmax(numbers ...int) int {
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}
func main() {
	defer handlePanic()
	max := Findmax(8, -5, -67, -12, 54)
	fmt.Println(max)
}
