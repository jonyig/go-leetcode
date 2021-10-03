package algorithm

import "log"

// 選擇排序
// O(n²)
// 找出最小值，往左邊為排序的值互換
func SelectionSort(list []int) []int {
	var (
		lowCarry int
		count    int
	)

	length := len(list)
	for i := 0; i <= length-1; i++ {
		lowCarry = i
		for j := i + 1; j <= length-1; j++ {
			count++
			if list[lowCarry] > list[j] {
				lowCarry = j
			}

		}
		list[i], list[lowCarry] = list[lowCarry], list[i]
	}
	log.Print(count)
	return list
}
