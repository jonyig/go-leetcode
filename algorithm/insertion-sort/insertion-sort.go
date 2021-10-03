package algorithm

import "log"

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
