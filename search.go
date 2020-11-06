package interview

import "sort"

// LinearSearch ...
func LinearSearch(data []int, target int) bool {

	for i := 0; i < len(data); i++ {
		if data[i] == target {
			return true
		}
	}

	return false
}

// LinearSearchRecursive ...
func LinearSearchRecursive(data []int, target int) bool {

	low := 0
	high := len(data)

	return linearSearchRecursive(data, target, low, high)

}

// linearSearchRecursive ...
func linearSearchRecursive(data []int, target, low, high int) bool {

	if target < low || target > high || low > high {
		return false
	}

	if data[low] == target {
		return true
	}
	return linearSearchRecursive(data, target, low+1, high)
}

// BinarySearch ...
func BinarySearch(data []int, target int) bool {

	low := 0
	high := len(data)
	sort.Ints(data) // neccessary

	if target > high || target < low {
		return false
	}

	for low <= high {
		mid := low + (high-low)/2
		// fmt.Printf("target = %v, low = %v, high = %v, mid = %v, data[mid] = %v\n", target, low, high, mid, data[mid])

		if data[mid] == target {
			return true
		} else if data[mid] < target {
			low = mid
		} else {
			high = mid
		}
	}

	return false
}

// BinarySearchRecursive ...
func BinarySearchRecursive(data []int, target int) bool {

	low := 0
	high := len(data)
	sort.Ints(data) // neccessary

	return binarySearchRecursive(data, target, low, high)

}

// binarySearchRecursive ...
func binarySearchRecursive(data []int, target, low, high int) bool {

	if target < low || target > high || low > high {
		return false
	}

	mid := low + (high-low)/2
	//fmt.Printf("target = %v, low = %v, high = %v, mid = %v, data[mid] = %v\n", target, low, high, mid, data[mid])

	if data[mid] == target {
		return true
	} else if data[mid] < target {
		return binarySearchRecursive(data, target, mid, high)
	} else {
		return binarySearchRecursive(data, target, low, mid)
	}
}
