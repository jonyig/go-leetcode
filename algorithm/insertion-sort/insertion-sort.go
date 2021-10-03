package algorithm

import "log"

// 插入排序
// O(n²)
// 依序讀取未排序過的數字，並且往前面找出可以插入的地方
func InsertionSort(list []int) []int {
	var (
		tempCarry   int
		targetCarry int
		count       int
	)
	length := len(list)
	for i := 1; i <= length-1; i++ {
		tempCarry = i - 1
		targetCarry = i

		for {
			if tempCarry < 0 {
				break
			}
			count++
			if list[targetCarry] < list[tempCarry] {
				list[targetCarry], list[tempCarry] = list[tempCarry], list[targetCarry]
				targetCarry = tempCarry
				tempCarry--
			} else {
				break
			}
		}
	}
	log.Print(count)

	return list
}
