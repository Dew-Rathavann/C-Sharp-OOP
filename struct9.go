package main

import "fmt"

func hello(t interface{}) {
	fmt.Printf("Hello %T\n", t)
}
func hi(t ...interface{}) {
	fmt.Print("Hi ")
	for _, v := range t {
		fmt.Printf("%T,", v)
	}
	fmt.Println()
}
func main() {
	hello(3.1415)
	hello(true)
	hello("Myself")
	hi("Mike", false, 10, 10e15)
}
