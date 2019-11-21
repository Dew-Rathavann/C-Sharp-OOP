package main

import "fmt"

func main() {
	fmt.Print("Input decimal:")
	var dec int
	fmt.Scan(&dec)
	fmt.Printf("Binary=%b", dec)
}
