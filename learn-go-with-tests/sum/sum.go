package main

func Sum(nums []int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
