package main

import "fmt"

func main() {
	fmt.Print("Input your name:")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello ", name)

}
