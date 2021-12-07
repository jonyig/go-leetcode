package leetcode

func searchInsert(nums []int, target int) int {
	ln := len(nums)
	i := 0
	for i < ln {

		if nums[i] > target {
			return i
		}
		if nums[i] == target {
			return i
		}

		i++
	}

	return i
}
