package ordering

func ReverseSlice(numbers []int) []int {
	size := len(numbers)

	for i := 0; i < size/2; i++ {
		numbers[i], numbers[size-1-i] = numbers[size-1-i], numbers[i]
	}

	return numbers
}
