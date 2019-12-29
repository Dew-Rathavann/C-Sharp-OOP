package main

import "fmt"

type Inter interface{}

func desc(i Inter) {
	fmt.Printf("%v,%T\n", i, i)
}
func main() {
	var i Inter
	i = 3.1415
	desc(i)
	i = "Goodbye"
	desc(i)
}
