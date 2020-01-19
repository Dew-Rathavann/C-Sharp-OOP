package main

import "fmt"

func sum(c chan int, number ...int) {
	sum := 0
	for _, v := range number {
		sum = sum + v
	}
	c <- sum
}
func printer(c chan int) {
	number := <-c
	fmt.Println(number)
}
func main() {
	c := make(chan int)
	go printer(c)
	go printer(c)
	go sum(c, 2, 4, 6)
	go sum(c, 12, 10)
	var input string
	fmt.Scan(&input)
}
