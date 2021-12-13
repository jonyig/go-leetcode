package leetcode

func maxSubArray(nums []int) int {
	result := nums[0]
	temp := nums[0]

	for _,item := range nums[1:]{
		temp = max(temp + item,item)
		result = max(result,temp)
	}
	return result
}

func max(a int ,b int ) int {
	if a > b {
		return a
	}

	return b
}
