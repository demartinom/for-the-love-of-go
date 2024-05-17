package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d given %v, want %d", got, numbers, want)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d given %v, want %d", got, numbers, want)
		}
	})
}

func ExampleSum() {
	numbers := []int{10, 20, 30}
	fmt.Println(Sum(numbers))
	// Output: 60
}

// func TestSumAll(t *testing.T) {
// 	got := SumAll([]int{1, 2}, []int{0, 9})
// 	want := []int{3, 9}
// 	if !slices.Equal(got, want) {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
