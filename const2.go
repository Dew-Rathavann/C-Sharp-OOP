package main

import "fmt"

func main() {
	for {
		fmt.Println("Before")
		break
		fmt.Println("After")
	}
	fmt.Println("Next statement")
	for i := 0; i < 3; i++ {
		fmt.Println("Before")
		continue
		fmt.Println("After")
	}
	fmt.Println("Next statement")
}
