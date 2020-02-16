package main
import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}
func consume(data chan int, done chan bool){
	
}
func main() {

}
