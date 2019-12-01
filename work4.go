package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Input: ")
	var text, a string
	a = "Hello World"
	fmt.Scan(&text)
	strings.Compare(text, "a")
	fmt.Println(a)
}
