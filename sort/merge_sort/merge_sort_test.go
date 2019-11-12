package merge_sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	var data []int = []int{6, 8, 45, 98, 543, -56, -9987, -67, 457, 36, 867}
	Sort(data)
	fmt.Println(data)
}
