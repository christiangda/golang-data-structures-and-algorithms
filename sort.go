package interview

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
