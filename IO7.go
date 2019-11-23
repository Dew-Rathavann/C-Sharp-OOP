package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Input your name and age: ")
	n, err := fmt.Scanln(&name, &age)
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Number of items successfully scaned: ", n)
	fmt.Println("Error", err)
}
