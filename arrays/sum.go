package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums;
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		// check that slice is not empty
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// sum the numbers in the slice, excluding the first number
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums;
}
