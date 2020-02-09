package main

import "fmt"

func main() {

	Map := make(map[string]int)

	Map["k1"] = 7
	Map["k2"] = 13

	fmt.Println("map:", Map)

	v1 := Map["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(Map))

	delete(Map, "k2")
	fmt.Println("map:", Map)

	_, prs := Map["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
