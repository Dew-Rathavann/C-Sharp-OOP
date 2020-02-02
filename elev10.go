package main

import (
	"fmt"
	"time"
)

var test = "testing"
var currtime = "15:04:05"
var date = "02/01/2006"

func main() {
	t := time.Now()
	date := t.Format("02/01/2006")
	currtime := t.Format("15:04:05")

	fmt.Println(test)
	fmt.Println(currtime)
	fmt.Println(date)
}
