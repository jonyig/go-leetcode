package algorithm

import "log"

// 氣泡排序
// O(n²)
// 依序比較兩個數字，前者大於後者就換位置
func BubbleSort(list []int) []int {
	length := len(list)
	var count int

	for {
		length--
		if length <= 0 {
			break
		}
		index := 0
		for {
			if index >= length {
				break
			}
			if list[index] > list[index+1] {
				list[index], list[index+1] = list[index+1], list[index]
				count++
			}
			index++
		}
	}
	log.Print(count)
	return list
}
