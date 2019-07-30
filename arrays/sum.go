package arrays

func sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func sumAll(numbers ...[]int) []int {
	var sums []int

	for _, number := range numbers {
		sums = append(sums, sum(number))
	}

	return sums
}

func sumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, sum(tail))
		}
	}

	return sums
}
