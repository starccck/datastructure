package quick_sort

func quickSort(data []int, start, end int) {
	if start > end {
		return
	}
	var i, j, x int
	x = data[start]
	for i , j = start, end; i < j; {
		for ; i < j && data[j] > x; j-- {}
		if i < j {
			data[i] = data[j]
			i++
		}


		for ; i < j && data[i] <= x; i++ {}
		if i < j {
			data[j] = data[i]
			j--
		}

	}
	data[i] = x

	quickSort(data, start, i - 1)
	quickSort(data, i + 1, end)
}

func quickSort2(data []int, start, end int) {
	if start >= end {
		return
	}

	var i, j int
	for i, j = start + 1, end; i <= j; {
		for ;i <= j && data[i] < data[start]; i++ {}
		for ; i <= j && data[j] >= data[start]; j-- {}

		if i <= j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[start], data[j] = data[j], data[start]

	quickSort2(data, start, j - 1)
	quickSort2(data, j + 1, end)
}