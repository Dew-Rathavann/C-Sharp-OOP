package main

import "fmt"

func main() {
	fmt.Printf("%v\n", 20)
	fmt.Printf("%#v\n", "6")
	fmt.Printf("%t\n", 18 <= 18)
	fmt.Printf("%s\n", "Dragon")
	fmt.Printf("%q\n", "#Hello!")
	fmt.Printf("%q\n", 35)
	fmt.Printf("%b\n", 26)
	fmt.Printf("%d\n", 26)
	fmt.Printf("%o\n", 26)
	fmt.Printf("%x\n", 26)
	fmt.Printf("%X\n", 26)
	fmt.Printf("%c\n", 35)
	fmt.Printf("%e\n", 0.1)
	fmt.Printf("%E\n", 0.1)
	fmt.Printf("%f\n", 12.1)
	fmt.Printf("%F\n", 12.1)
	fmt.Printf("%g\n", 12.1)
	fmt.Printf("%06d\n", 180)
}
