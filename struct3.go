package main

import (
	"bufio"
	"fmt"
	"os"
)

type student struct {
	name    string
	age     int
	major   string
	contact string
}

func (stu student) introduce() {
	fmt.Println("Your name is", stu.name)
	fmt.Println("You're", stu.age, "years old")
	fmt.Println("You're studying", stu.major)
	fmt.Println("Contact(phone):", stu.contact)
}

type add struct {
	address string
	student
}

func main() {
	var info add
	fmt.Print("Enter your name: ")
	fmt.Scan(&info.name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&info.age)
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your major: ")
	//info.major, _ = in.ReadString('\n')
	in.ReadString(' ')
	fmt.Scanln(&info.major)
	//line, err := in.ReadString('\n')
	fmt.Print("Enter your phone number: ")
	fmt.Scan(&info.contact)
	info.introduce()
}
