package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 99
	fmt.Println(x)
	fmt.Println(len(x))
}
