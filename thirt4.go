package main
import (
	"fmt"
	"strings"
)
func main() {
	var text string
	fmt.Print("Enter text: ")
	fmt.Scan(&text)
	for i := 0; i < strings.Count(text); i++ {
		if i%2 == 0 {
			fmt.Sprint(strings.ToUpper(text[i]))
		} else {
			fmt.Sprint(strings.ToLower(text[i]))
		}
	}
	fmt.Println(text)
}
