package ordering

func BubbleSort(numbers []int) []int {
	size := len(numbers)

	for i := 0; i < size; i++ {
		swapped := false

		for j := 0; j < size-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return numbers
}
