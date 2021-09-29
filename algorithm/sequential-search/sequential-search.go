package algorithm

func SequentialSearch(list []string, target string) int {
	for key, item := range list {
		if item == target {
			return key
		}
	}
	return 0
}
