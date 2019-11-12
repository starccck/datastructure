package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickTest(t *testing.T) {
	var data []int = []int{6, 8, 45, 98, 543, -56, -9987, -67, 457, 36, 867}
	quickSort(data, 0, len(data) - 1)
	fmt.Println(data)
}

func TestQuickSort2(t *testing.T) {
	var data []int = []int{6, 8, 45, 98, 543, -56, -9987, -67, 457, 36, 867}
	quickSort2(data, 0, len(data) - 1)
	fmt.Println(data)
}