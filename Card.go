package main

import "fmt"

var deck [52]card

type card struct {
	num  string
	suit string
}

func main() {
	var input int
Redo:
	fmt.Scan(&input)
	fmt.Print(pick(input))
	if input != 0 {
		goto Redo
	}
}
func pick(want int) string {
	for want < 0 {
		want = want + 52
	}
	for want > 52 {
		want = want - 52
	}
	num := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	suit := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	return fmt.Sprintf("%v %v", num[want%13], suit[int(want/13)])
}
