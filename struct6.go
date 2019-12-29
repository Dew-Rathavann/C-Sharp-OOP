package main

import "fmt"

type I interface {
	F()
}

func desc(i I) {
	fmt.Printf("%v,%T\n", i, i)
}

type T1 struct {
	text string
}

func (t T1) F() {
	fmt.Println(t.text)
}

type T2 struct {
	number int
}

func (t T2) F() {
	fmt.Println(t.number)
}
func main() {
	i1 := T1{"Go programming"}
	i2 := T2{3000}
	desc(i1)
	desc(i2)
}
