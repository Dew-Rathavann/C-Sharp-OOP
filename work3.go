package main

import "fmt"

func main() {
	fmt.Print("Input:")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Printf("a/b=%d\n", a/b)
	fmt.Println(a/b > c)
}
