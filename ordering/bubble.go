package ordering

func BubbleSort(numbers []int) []int {
	size := len(numbers)
	counter := size

	for i := 0; i < size; i++ {
		swapped := false
		for j := 0; j < counter; j++ {
			if (counter > j+1) && numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}

		counter -= 1

		if !swapped {
			break
		}
	}

	return numbers
}
