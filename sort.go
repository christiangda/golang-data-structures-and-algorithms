package interview

import (
	"math"
)

// BubbleSort ...
func BubbleSort(data []int) {

	max := len(data)

	for max > 0 {
		for i := 0; i < max-1; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
			}
		}
		max--
	}
}

// BubbleSortRecursive ...
func BubbleSortRecursive(data []int) {

	i := 0
	max := len(data)

	bubbleSortRecursive(data, i, max)
}

func bubbleSortRecursive(data []int, i, max int) {

	//fmt.Printf("data = %v, i = %v, max = %v\n", data, i, max)

	if i == max-1 {
		max--
		i = 0
	}

	if i < max-1 {
		if data[i] > data[i+1] {
			data[i], data[i+1] = data[i+1], data[i]
		}
		bubbleSortRecursive(data, i+1, max)
	}
}

// SelectionSort ...
func SelectionSort(data []int) {

	minIdx := 0
	minVal := math.MaxInt32

	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if data[j] < minVal {
				minIdx = j
				minVal = data[j]
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
		minVal = math.MaxInt32
	}
}
