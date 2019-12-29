package main

import "fmt"

type Inter interface{}

func main() {
	var i Inter
	i = "Hello"
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	b := i.(bool)
	fmt.Println(b)
}
