package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{2, 2, 4}, {1, 1, 2}, {5, 0, 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{{4, 2, 2}, {10, 8, 2}, {100, 40, 60}}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{{2, 2, 4}, {4, 8, 32}, {10, 10, 100}}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply (%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{{2, 2, 1}, {-1, -1, 1}, {10, 2, 5}}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("Divide (%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		a    float64
		want float64
	}{{4, 2}, {100, 10}}
	for _, tC := range testCases {
		got, err := calculator.Sqrt(tC.a)
		if err != nil {
			t.Fatal("Wrong!")
		}
		if tC.want != got {
			t.Errorf("want %f, got %f", tC.want, got)
		}
	}
}
