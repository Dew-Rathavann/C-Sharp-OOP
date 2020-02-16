package main
import (
	"fmt"
	"strings"
)
func main() {
	name := "Hello! Welcome home."
	words := strings.Split(name, "")
	fmt.Printf("%v \n", words)
}
