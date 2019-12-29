package main

import "fmt"

type teacher struct {
	name  string
	age   int
	sub   string
	tel   string
	email string
}

func (std teacher) introduce() {
	fmt.Println("The teacher's name is: ", std.name)
	fmt.Println("He's ", std.age)
	fmt.Println("He teaches ", std.sub)
	fmt.Println("Phone number: ", std.tel)
	fmt.Println("Email: ", std.email)
}
func main() {
	std := teacher{name: "Axle Blaze", age: 25, sub: "Programming language", tel: "0934416558", email: "AlexZabel@raymond.com"}
	std.introduce()
}
