package sort

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	arr := []int{1, 10, 3, 2, 6, 100, 88}
	QuickSort(arr)
	fmt.Println(arr)
}
