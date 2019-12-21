package Max

func Findmax(numbers ...int) int {
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}
