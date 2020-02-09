package main

import (
	"fmt"
	"math"
)

var a, b, c, dis float64

func CheckDis() {
	if dis > 0 {
		DisBiggerThan0()
	} else if dis < 0 {
		DisSmallerThan0()
	} else {
		DisEquals0()
	}
}
func DisBiggerThan0() (float64, float64) {
	fmt.Println("The equation has two different answers:")
	x1 := (-b - math.Sqrt(dis)) / 2 * a
	x2 := (-b + math.Sqrt(dis)) / 2 * a
	fmt.Println("x1= ", x1, "x2= ", x2)
	return x1, x2
}
func DisSmallerThan0() {
	fmt.Println("The equation has two different COMPLEX answers!")
}
func DisEquals0() float64 {
	fmt.Println("The equation has only one answer:")
	x1 := -b / 2 * a
	fmt.Println("x1= x2= ", x1)
	return x1
}
func main() {
	fmt.Println("The equation stays in the form of (ax^2+bx+c=0)")
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Print("Enter c: ")
	fmt.Scan(&c)
	dis = math.Pow(b, 2) - (4 * a * c)
	CheckDis()
}
