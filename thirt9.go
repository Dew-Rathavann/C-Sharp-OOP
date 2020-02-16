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
	s1 := student{
		firstName: "Fideo",
		lastName:  "Aldena",
		grade:     "A",
		country:   "Italy",
	}
	s2 := student{
		firstName: "Mark",
		lastName:  "Evans",
		grade:     "B",
		country:   "USA",
	}
	
}
