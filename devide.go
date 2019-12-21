package main

import "fmt"

func Devide(a int, b int) int {
	c := a / b
	return c
}
func HandlePanic() {
	text := recover()
	fmt.Println(text)
}
func main() {
	var a, b int
	fmt.Print("Input a: ")
	fmt.Scan(&a)
	fmt.Print("Input b: ")
	fmt.Scan(&b)
	fmt.Println(Devide(a, b))
	defer HandlePanic()
}
