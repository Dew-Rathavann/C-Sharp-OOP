package main

import "fmt"

type student struct {
	name string
	age  int
	id   string
}

func main() {
	var std [10]student
	std[0] = student{"Mike", 19, "61162110398-4"}
	fmt.Println(std[0])
	fmt.Println(std[0].name)
	fmt.Println(std[0].age)
	fmt.Println(std[0].id)
}
