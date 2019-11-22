package main

import "fmt"

func main() {
	fmt.Print("Input A:")
	var a float32
	fmt.Scan(&a)
	fmt.Print("Input B:")
	var b float32
	fmt.Scan(&b)
	fmt.Println("A+B=", a+b)
	fmt.Println("A-B=", a-b)
	fmt.Println("A*B=", a*b)
	fmt.Println("A/B=", a/b)

}
