package main

import (
	"fmt"
	"sync"
	"time"
)

func word(txt string, sleep time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(txt)
	time.Sleep(time.Millisecond * sleep)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go word("Hello", 2, &wg)
	go word("Hey", 1, &wg)
	wg.Wait()
	fmt.Println("GoodBye")
}
