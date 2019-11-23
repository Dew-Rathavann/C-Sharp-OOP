package main

import "fmt"

func main() {
	fmt.Print("Enter your name, gender, age and height: ")
	var name string
	var gen string
	var age int
	var height float32
	n, err := fmt.Scan(&name, &gen, &age, &height)
	fmt.Println(name, gen, age, height)
	fmt.Println(`Number of object: `, n)
	fmt.Println(`Number of error: `, err)
}
