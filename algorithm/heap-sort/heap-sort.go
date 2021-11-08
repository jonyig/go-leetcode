package algorithm

func HeapSortDesc(list []int) []int {
	heapSortInner(list, heapSortDescShiftDown)
	return list
}

func HeapSort(list []int) []int {
	heapSortInner(list, heapSortDescShift)
	return list
}
func heapSortInner(list []int, shiftDown func([]int, int, int)) {
	length := len(list)
	lengthDec := length - 1

	for start := length >> 1; start >= 0; start-- {
		shiftDown(list, start, lengthDec)
	}

	for i := lengthDec; i >= 1; i -= 1 {
		list[0], list[i] = list[i], list[0]
		shiftDown(list, 0, i-1)
	}
}
func heapSortDescShift(array []int, start int, end int) {
	for {
		left := (start << 1) + 1
		if left > end {
			break
		}

		right := left + 1

		var minChild int
		if right <= end && array[right] > array[left] {
			minChild = right
		} else {
			minChild = left
		}

		if array[start] >= array[minChild] {
			break
		}

		array[start], array[minChild] = array[minChild], array[start]
		start = minChild
	}
}

//[5, 10, 2, 7, 1]             5
//							10    2
//						  7	   1
// 0 ,1  ,2 ,3 ,4
func heapSortDescShiftDown(array []int, start int, end int) {
	for {
		left := (start << 1) + 1
		if left > end {
			break
		}

		right := left + 1
		var minChild int

		if right <= end && array[right] < array[left] {
			minChild = right
		} else {
			minChild = left
		}

		if array[start] <= array[minChild] {
			break
		}

		array[start], array[minChild] = array[minChild], array[start]
		start = minChild
	}
}
