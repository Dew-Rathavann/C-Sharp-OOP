package swap

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	x, y := swap("Hello ", "World")
	fmt.Println(x, y)
}
