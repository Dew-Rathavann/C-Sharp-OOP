package main

import (
	"fmt"
	"sort"
)

func main() {
	in := []int{3, 2, 1, 4, 3, 2, 1, 4, 1}
	sort.Ints(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		in[j] = in[i]
	}
	result := in[:j+1]
	fmt.Println(result)
}
