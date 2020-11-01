package interview

import (
	"testing"
)

func TestFibonacci(t *testing.T) {

	t.Run("Fibonacci iterative", func(t *testing.T) {
		tests := []struct {
			name  string
			given int64
			want  int64
		}{
			{name: "test case 1", given: 0, want: 0},
			{name: "test case 2", given: 1, want: 1},
			{name: "test case 3", given: 2, want: 1},
			{name: "test case 4", given: 3, want: 2},
			{name: "test case 5", given: 4, want: 3},
			{name: "test case 6", given: 5, want: 5},
			{name: "test case 7", given: 6, want: 8},
			{name: "test case 8", given: 7, want: 13},
			{name: "test case 9", given: 8, want: 21},
			{name: "test case 10", given: 9, want: 34},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				got := Fibonacci(tt.given)

				if got != tt.want {
					t.Errorf("Given Fibonacci(%v), got %v, want %v", tt.given, got, tt.want)
				}

			})
		}
	})

	t.Run("Fibonacci recursive", func(t *testing.T) {
		tests := []struct {
			name  string
			given int64
			want  int64
		}{
			{name: "test case 1", given: 0, want: 0},
			{name: "test case 2", given: 1, want: 1},
			{name: "test case 3", given: 2, want: 1},
			{name: "test case 4", given: 3, want: 2},
			{name: "test case 5", given: 4, want: 3},
			{name: "test case 6", given: 5, want: 5},
			{name: "test case 7", given: 6, want: 8},
			{name: "test case 8", given: 7, want: 13},
			{name: "test case 9", given: 8, want: 21},
			{name: "test case 10", given: 9, want: 34},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				got := FibonacciRecursive(tt.given)

				if got != tt.want {
					t.Errorf("Given Fibonacci(%v), got %v, want %v", tt.given, got, tt.want)
				}

			})

		}
	})
}
