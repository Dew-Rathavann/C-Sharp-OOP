package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("Computer Science", "e"))
	fmt.Println(strings.LastIndex("Computer Science", "e"))
	fmt.Println(strings.HasPrefix("Computer Science", "Computer"))
	fmt.Println(strings.HasSuffix("Computer Science", "Computer"))
	fmt.Println(strings.HasPrefix("Computer Science", "Science"))
	fmt.Println(strings.HasSuffix("Computer Science", "Science"))
}
