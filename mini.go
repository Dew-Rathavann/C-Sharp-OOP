package main

import "fmt"

var n int

func Findmini(number []int) int {
	mini := number[0]
	for i := 0; i < n; i++ {
		if number[i] < mini {
			mini = number[i]
		}
	}
	return mini
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
	fmt.Println(Findmini(num))
	defer HandlePanic()
}
