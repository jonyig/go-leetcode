package test0001

//func main() {
//	nums := []int{2, 7, 11, 15}
//
//	fmt.Print(twoSum(nums, 26))
//}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}

func test() int {
	return 123
}
