package main

import "fmt"

func main() {
	r := []bool{true, false, false, true}
	for _, v := range r {
		fmt.Printf("%t\n", v)
	}
}
