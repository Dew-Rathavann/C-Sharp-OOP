package main

import "fmt"

func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}
func main() {
	var n int
	fmt.Print("Input factorial you need to find: ")
	fmt.Scan(&n)
	fac := Factorial(n)
	fmt.Println(fac)
}
