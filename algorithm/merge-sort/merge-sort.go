package algorithm

import "log"

var count = 0

// 合併排序法
// O(n log n)
// 將所有值切成一個小個單位的區塊，並且依序合併起來
func MergeSort(list []int) []int {
	result := execute(list)
	log.Print(count)

	return result
}

func execute(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	mid := len(list) >> 1
	leftList := execute(list[:mid])
	rightList := execute(list[mid:])

	leftIndex := 0
	rightIndex := 0
	var (
		mergeList []int
	)

	for {
		if len(mergeList) >= len(list) {
			break
		}
		count++
		if rightIndex > len(rightList)-1 {
			mergeList = append(mergeList, leftList[leftIndex])
			leftIndex++
			continue
		} else if leftIndex > len(leftList)-1 {
			mergeList = append(mergeList, rightList[rightIndex])
			rightIndex++
			continue
		}

		if leftList[leftIndex] > rightList[rightIndex] {
			mergeList = append(mergeList, rightList[rightIndex])
			rightIndex++
			continue
		}

		if leftList[leftIndex] < rightList[rightIndex] {
			mergeList = append(mergeList, leftList[leftIndex])
			leftIndex++
			continue
		}
	}
	return mergeList
}
