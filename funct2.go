package main

import "fmt"

var x, y, z int

func sort() {
	if x > y && x > z {
		if y > z {
			fmt.Println(x, y, z)
		} else {
			fmt.Println(x, z, y)
		}
	} else if y > x && y > z {
		if x > z {
			fmt.Println(y, x, z)
		} else {
			fmt.Println(y, z, x)
		}
	} else if z > x && z > y {
		if x > y {
			fmt.Println(z, x, y)
		} else {
			fmt.Println(z, y, x)
		}
	}
}
func main() {
	fmt.Println("This program will sort three numbers from biggest to smallest.")
	fmt.Print("Enter x: ")
	fmt.Scan(&x)
	fmt.Print("Enter y: ")
	fmt.Scan(&y)
	fmt.Print("Enter z: ")
	fmt.Scan(&z)
	sort()
}
