package main

import (
	"fmt"
	"math"
)

func main() {
	var deg float64
	fmt.Print("Please enter degree you need to find: ")
	fmt.Scan(&deg)
	sin, cos := math.Sincos(deg)
	fmt.Printf("Sin= %.2f, Cos= %.2f", sin, cos)
}
