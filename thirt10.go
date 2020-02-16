package main
import "fmt"
func main() {
  numbers := []int{3, 12, 9, 7, 21, 31, 4, 10, 74}
  printSlice(numbers)
  fmt.Println("numbers ==", numbers)
  fmt.Println("numbers[1:4] ==", numbers[1:4])
  fmt.Println("numbers[:3] ==", numbers[:3])
}
