package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func (obj *Employee) Info() {
	if obj.Name == "" {
		obj.Name = "Rathavann"
	}
	if obj.Age == 0 {
		obj.Age = 18
	}
}

func main() {
	emp1 := Employee{Name: "Mr. Dew"}
	emp1.Info()
	fmt.Println(emp1)

	emp2 := Employee{Age: 19}
	emp2.Info()
	fmt.Println(emp2)
}
