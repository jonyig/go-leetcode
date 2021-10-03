package algorithm

//二元搜尋法
// O(log n) log 8 => 3
// 切半比較後，縮短大小，取得範圍
func BinarySearch(list []int, target int) int {
	start := 0
	end := len(list) - 1

	var (
		mid   int
		count int
	)

	for {
		if start > end {
			break
		}
		count++
		mid = (start + end) >> 1

		if list[mid] == target {
			return mid
		} else if list[mid] > target {
			end = mid + 1
		} else {
			start = mid - 1
		}
	}

	return -1
}
