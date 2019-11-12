package merge_sort

func Sort(data []int) {
	mergeSort(data, 0, len(data) - 1)
}

func mergeSort(data []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(data, start, mid)
	mergeSort(data, mid + 1, end)

	arr := make([]int, end - start + 1)
	loc := 0

	var i, j int
	for i, j = start, mid + 1; i <= mid && j <= end;  {
		if data[i] <= data[j] {
			arr[loc] = data[i]
			loc++
			i++
		} else {
				arr[loc] = data[j]
				loc++
				j++
		}
	}

	for i <= mid {
		arr[loc] = data[i]
		loc++
		i++
	}

	for j <= end {
		arr[loc] = data[j]
		loc++
		j++
	}

	for i := start; i <= end; i++ {
		data[i] = arr[i - start]
	}
}