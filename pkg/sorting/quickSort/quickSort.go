package quickSort

import "math/rand"

//сложность n*n максб n*logn средняя
//память так же
func QuickSort(slice []int) []int {

	if len(slice) == 0 {
		return nil
	}

	if len(slice) == 1 {
		return slice
	}

	if len(slice) == 2 {
		if slice[0] <= slice[1] {
			return []int{slice[0], slice[1]}
		}
		return []int{slice[1], slice[0]}
	}

	pivot := rand.Intn(len(slice))
	left := make([]int, 0, len(slice)) //n
	right := make([]int, 0, len(slice))
	sortedSlice := make([]int, 0, len(slice))
	for _, element := range slice {
		if element < pivot {
			left = append(left, element)
		}
		if element > pivot {
			right = append(right, element)
		}
	}

	sortedSlice = append(sortedSlice, QuickSort(left)...) //log n
	sortedSlice = append(sortedSlice, pivot)
	sortedSlice = append(sortedSlice, QuickSort(right)...) // log n

	return sortedSlice
}
