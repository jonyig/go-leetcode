package algorithm

//循序搜尋
//遍歷所有的數比對 O(n)
func SequentialSearch(list []int, target int) int {
	var count int
	for key, item := range list {
		count++
		if item == target {
			//log.Print(count)
			return key
		}
	}
	return 0
}
