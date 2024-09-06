package suggest

import (
	"fmt"
	"testing"
)

func TestBuildFastRadixTree(t *testing.T) {
	radix := BuildRadixTree()

	payload1 := radix.Search("i")
	payload2 := radix.Search("iph")
	payload3 := radix.Search("iphone")
	fmt.Println(payload1, payload2, payload3)
}
