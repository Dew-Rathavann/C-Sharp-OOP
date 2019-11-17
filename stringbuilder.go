package main

import (
	"fmt"
	"strings"
)

func main() {
	var a strings.Builder
	a.WriteString("Rathavann")
	a.WriteString(" ")
	a.WriteString("Phally")
	fmt.Println(a.String())
}
