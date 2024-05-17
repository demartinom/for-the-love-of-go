package main

import (
	"fmt"
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
