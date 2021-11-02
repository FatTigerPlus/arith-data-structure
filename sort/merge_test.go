package sort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	arr := []int{1, 10, 3, 2, 6, 100, 88}
	MergeSort(arr)
	fmt.Println(arr)
}
