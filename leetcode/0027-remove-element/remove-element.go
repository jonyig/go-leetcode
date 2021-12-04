package leetcode

func removeElement(nums []int, val int) int {
	ln := len(nums)

	j := 0 // points to  the index of last filled position
	for i := 0; i < ln; i++ {
		if val != nums[i] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
