package main

import "fmt"

func main() {
	fmt.Print("Input number:")
	var x float32
	fmt.Scan(&x)
	fmt.Printf("%e", x)
}
