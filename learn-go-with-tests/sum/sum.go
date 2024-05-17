package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	lengthOfNumbers := len(nums)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range nums {
		sums[i] = Sum(numbers)
	}
	return sums
}
