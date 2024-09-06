package mergeSort

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var Counter int32

func NextPowerOfTwo(x int) int {
	powers := [64]int{}
	for i := 0; i < 64; i++ {
		powers[i] = 1 << i
	}

	for i, _ := range powers {
		if x == powers[i] {
			return x
		}
		if x < powers[i] {
			x = powers[i-1]
			break
		}
	}
	return x
}

func MergeSortRecursive(slice []int) []int {
	buff := make([]int, 0, len(slice))
	SortRecursive(slice, buff)
	return slice
}

func MergeSortParallel(slice []int) []int {
	buff := make([]int, 0, len(slice))
	cpus := runtime.NumCPU()
	chunkSize := len(slice) / cpus

	if chunkSize < 1024 {
		chunkSize = 1024
	}

	chunkSize = NextPowerOfTwo(chunkSize)

	SortRecursiveParallel(slice, buff, chunkSize)
	return slice
}

func SortRecursiveParallel(slice []int, buff []int, chunkSize int) {

	atomic.AddInt32(&Counter, 1)

	if len(slice) == 1 {
		return
	}

	if len(slice) == 2 {
		if slice[1] < slice[0] {
			temp := slice[0]
			slice[0] = slice[1]
			slice[1] = temp
		}
		return
	}

	mid := (len(slice)) >> 1

	if len(slice) > chunkSize && len(slice) < chunkSize*2 {
		buffLeft := make([]int, 0, len(slice)>>1)
		buffRight := make([]int, 0, len(slice)>>1)
		var wg sync.WaitGroup

		wg.Add(2)

		go func() {
			defer wg.Done()
			SortRecursiveParallel(slice[0:mid], buffLeft, chunkSize)
		}()

		go func() {
			defer wg.Done()
			SortRecursiveParallel(slice[mid:], buffRight, chunkSize)
		}()
		wg.Wait()

	} else {
		SortRecursiveParallel(slice[0:mid], buff, chunkSize)
		SortRecursiveParallel(slice[mid:], buff, chunkSize)
	}

	buff = Merge(slice[0:mid], slice[mid:], buff)
	copy(slice, buff)
}

func SortRecursive(slice []int, buff []int) {
	if len(slice) == 1 {
		return
	}

	if len(slice) == 2 {
		if slice[1] < slice[0] {
			temp := slice[0]
			slice[0] = slice[1]
			slice[1] = temp
		}
		return
	}
	mid := (len(slice)) >> 1

	SortRecursive(slice[0:mid], buff)
	SortRecursive(slice[mid:], buff)

	buff = Merge(slice[0:mid], slice[mid:], buff)
	copy(slice, buff)
}

func MergeSort(slice []int) []int {
	slice = PartialSort(slice)

	buff := make([]int, 0, len(slice)>>1)

	if len(slice) <= 2 {
		return slice
	}

	if len(slice) == 3 {
		left := slice[:1]
		right := slice[1:3]
		slice = Merge(left, right, buff)
		buff = buff[0:0]
		return slice
	}

	step := 2

	for step <= (len(slice) >> 1) {
		endRight := len(slice)
		for {
			startRight := endRight - step
			if startRight < 0 {
				break
			}
			if startRight == 0 {

			}

			endLeft := startRight
			startLeft := startRight - step

			if startLeft < 0 {
				break
			}

			left := slice[startLeft:endLeft]
			right := slice[startRight:endRight]

			buff = Merge(left, right, buff)
			j := 0
			for i := startLeft; i <= (endRight - 1); i++ {
				slice[i] = buff[j]
				j++
			}
			buff = buff[0:0]
			endRight = startLeft

			if endRight == 0 {
				break
			}
		}
		step = step * 2
	}
	return slice
}

func PartialSort(slice []int) []int {
	stack := NewStack(len(slice))

	stack.Push(0)
	stack.Push(len(slice))

	for {
		if stack.Size() == 0 {
			break
		}
		end := stack.Pop()
		start := stack.Pop()
		diff := end - start

		if diff == 1 {
			continue
		}

		if diff == 2 {
			if slice[start] > slice[end-1] {
				temp := slice[start]
				slice[start] = slice[end-1]
				slice[end-1] = temp
			}
			continue
		}

		mid := start + (diff >> 1)

		stack.Push(start)
		stack.Push(mid)

		stack.Push(mid)
		stack.Push(end)
	}
	return slice
}

func Merge(left []int, right []int, buff []int) []int {
	startLeft := 0
	startRight := 0
	buff = buff[0:0]

	for startLeft < len(left) && startRight < len(right) {

		if left[startLeft] < right[startRight] {
			buff = append(buff, left[startLeft])
			startLeft++
		} else {
			buff = append(buff, right[startRight])
			startRight++
		}
	}
	buff = append(buff, left[startLeft:]...)
	buff = append(buff, right[startRight:]...)
	return buff
}
