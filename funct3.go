package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total = total + n
	}
	return total
}
func main() {
	x := sum(3, 5, 4, 1, 7)
	fmt.Println(x)
	y := sum()
	fmt.Println(y)
}
