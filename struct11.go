package main

import "fmt"

type myError struct {
	error string
}

func (e myError) Error() string {
	return e.error
}
func say(word string) error {
	if word == "happy new year" {
		return myError{"Hooray!"}
	}
	return nil
}
func main() {
	e1 := say("hello")
	fmt.Println(e1)
	e2 := say("happy new year")
	fmt.Println(e2)
}
