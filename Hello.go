package Hello

import "strings"

func Find(x string) bool {
	if strings.Contains(x, "w") {
		return true
	}
	return false
}
