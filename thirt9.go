package main
import "fmt"
type student struct{
	firstName string
	lastName  string
	grade     string
	country   string
}
func filter(stu []student, f func(student) bool) []student{
	var r []student
	for _, v := range stu {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}
func main() {

}
