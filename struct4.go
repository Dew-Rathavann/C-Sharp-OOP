package main

import "fmt"

type People struct {
	name string
	age  int
}

func (grow *People) growUp(gr int) {
	grow.age = grow.age + gr
}
func main() {
	var a People
	a.age = 18
	fmt.Println(a.age)
	a.growUp(10)
	fmt.Println(a.age)
}
