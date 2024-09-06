package mergeSort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestParallelMergeSort(t *testing.T) {

	slice := make([]int, 10*1024*1024, 10*1024*1024)

	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(len(slice))
	}

	t.Run("simple", func(t *testing.T) {
		t.Skip()
		start := time.Now()

		MergeSortRecursive(slice)
		fmt.Println(time.Now().Sub(start))
	})

	t.Run("parallel", func(t *testing.T) {
		start := time.Now()

		MergeSortParallel(slice)
		fmt.Println(time.Now().Sub(start))
		fmt.Println("prev: 123840615")
		fmt.Println("counter: ", Counter)
	})

}

//func TestMergeSort(t *testing.T) {
//	t.Run("slice of three", func(t *testing.T) {
//		slice := []int{7, 5, 4}
//		MergeSort(slice)
//		assert.Equal(t, []int{4, 5, 7}, MergeSort(slice))
//	})
//
//	t.Run("slice of five", func(t *testing.T) {
//		slice := []int{7, 5, 4, 3, 2}
//		MergeSort(slice)
//		assert.Equal(t, []int{2, 3, 4, 5, 7}, MergeSort(slice))
//	})
//
//	t.Run("even size", func(t *testing.T) {
//		slice := []int{7, 5, 4, 2, 1, 3, 5, 8}
//		assert.Equal(t, []int{1, 2, 3, 4, 5, 5, 7, 8}, MergeSort(slice))
//	})
//}

func TestMergeSortRecursive(t *testing.T) {
	t.Run("slice of three", func(t *testing.T) {
		slice := []int{7, 5, 4}
		assert.Equal(t, []int{4, 5, 7}, MergeSortRecursive(slice))
	})

	t.Run("slice of five", func(t *testing.T) {
		slice := []int{7, 5, 4, 3, 2}
		assert.Equal(t, []int{2, 3, 4, 5, 7}, MergeSortRecursive(slice))
	})

	t.Run("even size", func(t *testing.T) {
		slice := []int{7, 5, 4, 2, 1, 3, 5, 8}
		assert.Equal(t, []int{1, 2, 3, 4, 5, 5, 7, 8}, MergeSortRecursive(slice))
	})
}

func TestMerge(t *testing.T) {
	t.Run("merge2+2", func(t *testing.T) {
		left := []int{1, 5}
		right := []int{2, 3}
		buff := make([]int, 0, 4)
		merged := Merge(left, right, buff)
		assert.Equal(t, []int{1, 2, 3, 5}, merged)
	})

	t.Run("merge1+2", func(t *testing.T) {
		left := []int{2}
		right := []int{1, 3}
		buff := make([]int, 0, 3)
		merged := Merge(left, right, buff)
		assert.Equal(t, []int{1, 2, 3}, merged)
	})

	t.Run("empty", func(t *testing.T) {
		left := []int{}
		right := []int{1, 3}
		buff := make([]int, 0, 2)
		merged := Merge(left, right, buff)
		assert.Equal(t, []int{1, 3}, merged)
	})
}
