package interview

import (
	"testing"
)

func TestFactorial(t *testing.T) {

	t.Run("Test Factorial iterativeve", func(t *testing.T) {
		tests := []struct {
			name  string
			given int64
			want  int64
		}{
			{name: "test case 1", given: -1, want: 0},
			{name: "test case 2", given: 1, want: 1},
			{name: "test case 3", given: 2, want: 2},
			{name: "test case 4", given: 3, want: 6},
			{name: "test case 5", given: 4, want: 24},
			{name: "test case 6", given: 5, want: 120},
			{name: "test case 7", given: 6, want: 720},
			{name: "test case 8", given: 7, want: 5040},
			{name: "test case 9", given: 8, want: 40320},
			{name: "test case 10", given: 9, want: 362880},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {
				got := Factorial(tt.given)

				if got != tt.want {
					t.Errorf("Given Factorial(%v), got %v, want %v", tt.given, got, tt.want)
				}

			})
		}
	})

	t.Run("Test Factorial recursive", func(t *testing.T) {
		tests := []struct {
			name  string
			given int64
			want  int64
		}{
			{name: "test case 1", given: -1, want: 0},
			{name: "test case 2", given: 1, want: 1},
			{name: "test case 3", given: 2, want: 2},
			{name: "test case 4", given: 3, want: 6},
			{name: "test case 5", given: 4, want: 24},
			{name: "test case 6", given: 5, want: 120},
			{name: "test case 7", given: 6, want: 720},
			{name: "test case 8", given: 7, want: 5040},
			{name: "test case 9", given: 8, want: 40320},
			{name: "test case 10", given: 9, want: 362880},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				got := FactorialRecursive(tt.given)

				if got != tt.want {
					t.Errorf("Given FactorialRecursive(%v) got %v, want %v", tt.given, got, tt.want)
				}
			})
		}
	})
}
