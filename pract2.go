package main

import "fmt"

func myfn1(i string) {
	fmt.Println(i)
}

func myfunc2(firstName string, lastName string) string {

	return "Hello " + firstName + " " + lastName + "!"

}

func test(do func(string), val string) {
	do(val)
}

func test1(t func(string, string) string, fname string, lname string) string {

	opt := t(fname, lname)
	return opt
}

func main() {
	test(myfn1, "Aishu")
	greet := test1(myfunc2, "Aishu", "S")
	fmt.Println(greet)

}
