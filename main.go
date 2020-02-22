package main
import (
	"fmt"
 	"os"
	"io/ioutil"
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
	return r
}
func FindFileFromExtension(extension []string, dir string, files *[]string){
	fs, err := ioutil.ReadDir(dir)
	if err == nil{
		for _, f := range fs{
			for _, ex := range extension{
				if strings.HasSuffix(f.Name(), ex){
					
				}
			}
		}
	}
}
func main() {

}
