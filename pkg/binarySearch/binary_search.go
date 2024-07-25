package binarySearch

func BinarySearch3(value int, slice []int) int {
	indexStart := 0
	indexEnd := len(slice) - 1
	indexMiddle := 0

	if value < slice[indexStart] {
		return 0
	}
	if value > slice[indexEnd] {
		return indexEnd + 1
	}

	for indexStart <= indexEnd {
		indexMiddle = (indexStart + indexEnd) >> 1

		if value == slice[indexMiddle] {
			return indexMiddle
		}

		if value < slice[indexMiddle] {
			indexEnd = indexMiddle - 1
		}

		if value > slice[indexMiddle] {
			indexStart = indexMiddle + 1
		}
	}
	return indexStart
}

func BinarySearch(value int, slice []int) bool {
	var middleIndex int
	res := false
	if len(slice) == 0 {
		return false
	}
	if len(slice) <= 2 {
		for _, element := range slice {
			if value == element {
				return true
			}
		}
		return false
	}

	// 00100010
	// 00010001
	middleIndex = len(slice) >> 1

	if value == slice[middleIndex] {
		res = true
		return res
	}
	if value < slice[middleIndex] {
		slice = slice[:middleIndex]
		return BinarySearch(value, slice)
	}
	slice = slice[middleIndex:]
	return BinarySearch(value, slice)
}

func BinarySearch2(value int, slice []int) int {
	indexStart := 0
	indexEnd := len(slice) - 1
	indexMiddle := len(slice) >> 1

	if value == slice[indexMiddle] {
		return indexMiddle
	}

	if value < slice[indexStart] {
		return 0
	}

	if value < slice[indexMiddle] {
		leftSlice := slice[:indexMiddle]
		return BinarySearch2(value, leftSlice)
	}

	if value > slice[indexEnd] {
		return indexEnd + 1
	}

	indexStart = indexMiddle
	rightSlice := slice[indexStart:]
	return indexStart + BinarySearch2(value, rightSlice)
}
