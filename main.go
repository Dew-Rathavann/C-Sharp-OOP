package main
import (
  "fmt"
  "os"
  "strings"
)
func getDrive() (r string){
  	var drive string
	fmt.Print("Please select your drive to search(C/D): ")
	fmt.Scan(&drive)
  	f, err := os.Open(string(drive) + ":\\")
	if err == nil {
		d := string(drive) + ":/"
		r = append(r, d)
		f.Close()
	}
	
}
func main() {

}
