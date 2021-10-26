package algorithm

func QuickSort(list []int) []int {
	sortRecursively(list, 0, len(list)-1)
	return list
}

//list := []int{69, 81, 30, 38, 9, 2, 47, 61, 32, 79}

func sortRecursively(list []int, left int, right int) {

	if left < right {

		pivot := list[(left+right)/2]
		l := left
		r := right

		for {

			for pivot > list[l] {
				l++
			}

			for pivot < list[r] {
				r--
			}

			if l >= r {
				break
			}
			list[r], list[l] = list[l], list[r]
		}
		sortRecursively(list, left, l-1)
		sortRecursively(list, r+1, right)
	}

}
