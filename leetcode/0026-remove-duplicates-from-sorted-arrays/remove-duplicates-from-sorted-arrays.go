package leetcode

func removeDuplicates(nums []int) int {
	point := -101
	var count int
	for count < len(nums) {
		if point < nums[count] {
			point = nums[count]
			count++
			continue
		}

		nums = append(remove(nums,count))
	}
	return count
}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func removeDuplicatesBest(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}

	j := 0 // points to  the index of last filled position
	for i := 1; i < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
