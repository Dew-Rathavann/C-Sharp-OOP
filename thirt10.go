package main
import "fmt"
func printSlice(x []int){
  
}
func main() {
  numbers := []int{3, 12, 9, 7, 21, 31, 4, 10, 74}
  printSlice(numbers)
  fmt.Println("numbers ==", numbers)
  fmt.Println("numbers[1:4] ==", numbers[1:4])
  fmt.Println("numbers[:3] ==", numbers[:3])
  fmt.Println("numbers[4:] ==", numbers[4:])
  numbers1 := make([]int, 0, 5)
  printSlice(numbers1)
  number2 := numbers[:2]
  printSlice(number2)
  number3 := numbers[2:5]
  printSlice(number3)
}
