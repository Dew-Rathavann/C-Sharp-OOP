package main

import "fmt"

var id, name, gen, grade []string
var score []int
var n, i int

func input() {

	id = make([]string, n)
	name = make([]string, n)
	gen = make([]string, n)
	grade = make([]string, n)
	score = make([]int, n)
	for i = 0; i < n; i++ {
		fmt.Print("Input ID: ")
		fmt.Scan(&id[i])
		fmt.Print("Input Name: ")
		fmt.Scan(&name[i])
		fmt.Print("Input Gender: ")
		fmt.Scan(&gen[i])
		fmt.Print("Input Score: ")
		fmt.Scan(&score[i])
		if score[i] >= 80 {
			grade[i] = "A"
		} else if score[i] < 80 && score[i] >= 70 {
			grade[i] = "B"
		} else if score[i] < 70 && score[i] >= 60 {
			grade[i] = "C"
		} else if score[i] < 60 && score[i] >= 50 {
			grade[i] = "D"
		} else {
			grade[i] = "F"
		}
	}
}
func output() {
	fmt.Printf("ID\tName\tGender\tScore\tGrade\n")
	for i = 0; i < n; i++ {
		fmt.Printf("%v\t%v\t%v\t%d\t%v\n", id[i], name[i], gen[i], score[i], grade[i])
	}
}
func main() {
	var option int
	for option != 3 {
		fmt.Println("1: Input Information")
		fmt.Println("2: Output Information")
		fmt.Println("3: Exit Program")
		fmt.Print("Please Select an option to continue: ")
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Print("Input number of students: ")
			fmt.Scan(&n)
			input()
			break
		case 2:
			output()
			break
		case 3:
			break

		}
	}
}
