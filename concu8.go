package main

import "fmt"

func Sum(c chan int, number ...int) {
	sum := 0
	for _, v := range number {
		sum = sum + v
	}
	c <- sum
}
func Printer(c1, c2 chan int) {
	select {
	case num1 := <-c1:
		fmt.Println("Channel-1: ", num1)
	case num2 := <-c2:
		fmt.Println("Channel-2: ", num2)
	}
}
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go Printer(c1, c2)
	go Sum(c1, 5, 6, 7)
	go Sum(c2, 20, 30)
	var input string
	fmt.Scan(&input)
}
