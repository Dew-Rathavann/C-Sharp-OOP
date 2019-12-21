package main

import "fmt"

var n int

func Findmax(number []int) int {
	max := number[0]
	for i := 0; i < n; i++ {
		if number[i] > max {
			max = number[i]
		}
	}
	return max
}

func HandlePanic() {
	text := recover()
	fmt.Println(text)
}
func main() {
	var num []int
	fmt.Print("Input number of index: ")
	fmt.Scan(&n)
	num = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("Input idex: ")
		fmt.Scan(&num[i])
	}
	fmt.Println(Findmax(num))
	defer HandlePanic()
}
