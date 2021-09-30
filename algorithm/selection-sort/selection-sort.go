package algorithm

func SelectionSort(list []int) []int {
	var (
		lowCarry int
	)

	length := len(list)
	for i := 0; i <= length-1; i++ {
		lowCarry = i
		for j := i + 1; j <= length-1; j++ {
			if list[lowCarry] > list[j] {
				lowCarry = j
			}

		}
		list[i], list[lowCarry] = list[lowCarry], list[i]
	}

	return list
}
