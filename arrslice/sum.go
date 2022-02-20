package arrslice

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arrs ...[]int) []int {
	var sums []int

	for _, numbers := range arrs {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(arrs ...[]int) []int {
	var sums []int
	for _, arr := range arrs {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			tail := arr[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
