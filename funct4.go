package main

import "fmt"

func language(greet string) func(string) string {
	return func(name string) string {
		return greet + name
	}
}
func main() {
	x := language("World")
	fmt.Println(("Hello "), x)
	fmt.Println(("Goodbye "), x)
}
