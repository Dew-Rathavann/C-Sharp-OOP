package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Welcome to my house", "house"))
	fmt.Println(strings.Contains("Welcome to my house", "House"))
	fmt.Println(strings.ContainsAny("Welcome to my house", "Me"))
	fmt.Println(strings.ContainsAny("Welcome to my house", "time"))
}
