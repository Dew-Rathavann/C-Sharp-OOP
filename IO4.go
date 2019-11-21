package main

import "fmt"

func main() {
	fmt.Print("Input ID:")
	var id int
	fmt.Scan(&id)
	fmt.Print("Input Score:")
	var score float32
	fmt.Scan(&score)
	fmt.Print("Input your Name:")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Datatype of ID:%T\n", id)
	fmt.Printf("Datatype of Score:%T\n", score)
	fmt.Printf("Datatype of Name:%T\n", name)
}
