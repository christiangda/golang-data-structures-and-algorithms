package interview

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	t.Run("iterative solution", func(t *testing.T) {
		tests := []struct {
			data  []int
			given int
			want  bool
		}{
			{data: data, given: 25, want: false},
			{data: data, given: -25, want: false},
			{data: data, given: 2, want: true},
			{data: data, given: 16, want: true},
			{data: data, given: 11, want: true},
			{data: data, given: 19, want: true},
			{data: data, given: 1, want: true},
		}

		for _, tt := range tests {
			got := LinearSearch(tt.data, tt.given)
			if got != tt.want {
				t.Errorf("LinearSearch(%v), got %v, want % v", tt.given, got, tt.want)
			}

		}
	})

	t.Run("recursive solution", func(t *testing.T) {
		tests := []struct {
			data  []int
			given int
			want  bool
		}{
			{data: data, given: 25, want: false},
			{data: data, given: -25, want: false},
			{data: data, given: 2, want: true},
			{data: data, given: 16, want: true},
			{data: data, given: 11, want: true},
			{data: data, given: 19, want: true},
			{data: data, given: 1, want: true},
		}

		for _, tt := range tests {
			got := LinearSearchRecursive(tt.data, tt.given)
			if got != tt.want {
				t.Errorf("LinearSearchRecursive(%v), got %v, want % v", tt.given, got, tt.want)
			}

		}
	})
}
func TestBinarySearch(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	t.Run("iterative solution", func(t *testing.T) {
		tests := []struct {
			name  string
			data  []int
			given int
			want  bool
		}{
			{name: "test case 1", data: data, given: 25, want: false},
			{name: "test case 2", data: data, given: -25, want: false},
			{name: "test case 3", data: data, given: 2, want: true},
			{name: "test case 4", data: data, given: 16, want: true},
			{name: "test case 5", data: data, given: 11, want: true},
			{name: "test case 6", data: data, given: 19, want: true},
			{name: "test case 7", data: data, given: 1, want: true},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				got := BinarySearch(tt.data, tt.given)
				if got != tt.want {
					t.Errorf("BinarySearch(%v), got %v, want % v", tt.given, got, tt.want)
				}
			})
		}
	})

	t.Run("recursive solution", func(t *testing.T) {
		tests := []struct {
			name  string
			data  []int
			given int
			want  bool
		}{
			{name: "test case 1", data: data, given: 25, want: false},
			{name: "test case 2", data: data, given: -25, want: false},
			{name: "test case 3", data: data, given: 2, want: true},
			{name: "test case 4", data: data, given: 16, want: true},
			{name: "test case 5", data: data, given: 11, want: true},
			{name: "test case 6", data: data, given: 19, want: true},
			{name: "test case 7", data: data, given: 1, want: true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := BinarySearchRecursive(tt.data, tt.given)
				if got != tt.want {
					t.Errorf("BinarySearchRecursive(%v), got %v, want % v", tt.given, got, tt.want)
				}
			})
		}
	})
}
