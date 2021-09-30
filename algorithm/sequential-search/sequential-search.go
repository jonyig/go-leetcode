package algorithm

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
