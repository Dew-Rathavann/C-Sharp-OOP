package main

import "fmt"

type Integer struct {
	error string
}

func (e myError) Error() string {
	return e.error
}
func Input(num int) error {
	if num > 0 {
		return Integer{num}
	}
	return Integer{"Error!"}
}
func main() {
	n1 := Input(22)
	fmt.Println(n1)
	n2 := Input(-4)
	fmt.Println(n2)
}
