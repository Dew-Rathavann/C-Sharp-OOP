package main
import (
	"fmt"
	"strings"
)
func main() {
	var txt strings.Builder
	txt.WriteString("Today is not ")
	txt.WriteString("my lucky day")
	fmt.Println(txt.String())
	fmt.Println(strings.Contains(txt.String(), "get"))
	fmt.Println(strings.Count(txt.String(), "a"))
}
