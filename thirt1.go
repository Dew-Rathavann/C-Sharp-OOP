package main

import "fmt"
func Today(channel chan bool){
	fmt.Println("Friday")
	channel<-true
}
func main() {
	
}
