package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("One and only", "and"))
	fmt.Println(strings.Index("One and only", "Only"))
}
