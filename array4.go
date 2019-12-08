package main

import "fmt"

func main() {
	alpha := [4]string{"A", "B", "C", "D"}
	fmt.Println(alpha)
	x := alpha[0:1]
	fmt.Println(x)
	y := alpha[1:3]
	fmt.Println(y)
	z := y[0:1]
	fmt.Println(z)
	z[0] = "H"
	fmt.Println(alpha)
}
