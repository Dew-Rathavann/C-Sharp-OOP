package main

import "fmt"

func main() {
	name := []string{"Alfa", "Belta", "Gamma"}
	other := []string{"Landa", "hela"}
	name = append(name, other...)
	fmt.Println(name)
}
