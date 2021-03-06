package arraysnslices

func Sum(slice []int) int {
	var sum int

	for _, number := range slice {
		sum += number
	}

	return sum
}

func SumAll(slices ...[]int) (result []int) {
	for _, numbers := range slices {
		result = append(result, Sum(numbers))
	}

	return
}

func SumSliceTail(slices ...[]int) (result []int) {
	for _, numbers := range slices {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			tail := numbers[1:]
			result = append(result, Sum(tail))
		}
	}
	return
}
