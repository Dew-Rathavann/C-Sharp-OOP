package main

import (
	"fmt"
	"time"
)

func speak(txt string) {
	for i := 0; i < 3; i++ {
		fmt.Println(time.Now(), ":", i, ":", txt)
		time.Sleep(time.Millisecond)
	}
}
func main() {
	go speak("Good morning!")
	go speak("How are you?")
	var input string
	fmt.Scanln(&input)
}
