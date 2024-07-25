package main

import (
	"fmt"

	"github.com/polina/grammar/pkg/binarySearch"
)

func main() {
	mySlice := []int{1, 3, 5}
	res := binarySearch.BinarySearch3(6, mySlice)
	fmt.Println(res)
}
