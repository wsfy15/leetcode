/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	stack := []int{}
	for i := 0; i < n * 2; i++ {
		index := i % n
		for len(stack) > 0 && nums[index] > nums[stack[len(stack) - 1]] {
			res[stack[len(stack) - 1]] = nums[index]
			stack = stack[:len(stack) - 1]
		} 
		stack = append(stack, index)
	}

	return res
}
// @lc code=end

