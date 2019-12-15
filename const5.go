package main

import "fmt"

func main() {
	var a int
Read:
	fmt.Print("type number not higher than 50: ")
	fmt.Scan(&a)
	if a < 50 {
		fmt.Println(a)
		goto Read
	} else {
		fmt.Println("The number you entered is invalid.")
	}

}
