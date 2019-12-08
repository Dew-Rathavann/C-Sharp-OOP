package main

import "fmt"

func main() {
	number := [5]int{2, 7, 6, 1, 9}
	fmt.Println(number)
	fmt.Println("The third index is: ", number[3])
	number[3] = 8
	fmt.Println("Change the third index from '1' to ", number[3])
	fmt.Println(number)
	length := len(number)
	fmt.Println("Length= ", length)
}
