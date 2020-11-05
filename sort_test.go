package interview

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	t.Run("iterative algorithm", func(t *testing.T) {
		tests := []struct {
			name  string
			given []int
			want  []int
		}{
			{
				name:  "test 1",
				given: []int{10, 18, 6, 2, 4, 16, 8, 14, 12},
				want:  []int{2, 4, 6, 8, 10, 12, 14, 16, 18},
			},
			{
				name:  "test 2",
				given: []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				got := make([]int, len(tt.given))
				copy(got, tt.given)
				BubbleSort(got)

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("BubbleSort(%v), got: %v, want: %v", tt.given, got, tt.want)
				}

			})
		}
	})

	t.Run("recursive algorithm", func(t *testing.T) {
		tests := []struct {
			name  string
			given []int
			want  []int
		}{
			{
				name:  "test 1",
				given: []int{10, 18, 6, 2, 4, 16, 8, 14, 12},
				want:  []int{2, 4, 6, 8, 10, 12, 14, 16, 18},
			},
			{
				name:  "test 2",
				given: []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				got := make([]int, len(tt.given))
				copy(got, tt.given)
				BubbleSortRecursive(got)

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("BubbleSortRecursive(%v), got: %v, want: %v", tt.given, got, tt.want)
				}

			})
		}
	})
}

func TestSelectionSort(t *testing.T) {

	t.Run("iterative function", func(t *testing.T) {

		tests := []struct {
			name  string
			given []int
			want  []int
		}{
			{
				name:  "test 1",
				given: []int{10, 18, 6, 2, 4, 16, 8, 14, 12},
				want:  []int{2, 4, 6, 8, 10, 12, 14, 16, 18},
			},
			{
				name:  "test 2",
				given: []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				got := make([]int, len(tt.given))
				copy(got, tt.given)

				SelectionSort(got)

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("SelectionSort(%v), got %v, want %v", tt.given, got, tt.want)
				}

			})

		}

	})
}
