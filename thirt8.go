package main
import "fmt"
func myfun(a interface{}){
  switch a.(type){
    case int:
		  fmt.Println("Type: int, Value:", a.(int))
    
  }
}
func main() {

}
