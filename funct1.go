package main

import "fmt"

var x, y, sum, sub, mult int

func Calculate() {
	sum = x + y
	sub = x - y
	mult = x * y
}
func main() {
	fmt.Print("Input x: ")
	fmt.Scan(&x)
	fmt.Print("Input y: ")
	fmt.Scan(&y)
	Calculate()
	fmt.Println("x+y= ", sum)
	fmt.Println("x-y= ", sub)
	fmt.Println("x*y= ", mult)
}
