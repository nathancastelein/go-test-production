package unit

import (
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		numbers  []int
		expected int
	}{
		{numbers: []int{1, 2, 3}, expected: 6},
		{numbers: []int{4, 5, 6}, expected: 15},
		{numbers: []int{}, expected: 0},
	}

	for _, tc := range testCases {
		got := Sum(tc.numbers...)
		if got != tc.expected {
			t.Errorf("Sum(%v) = %d; want %d", tc.numbers, got, tc.expected)
		}
	}
}

func Sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
