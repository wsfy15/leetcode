/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132模式
 */

// @lc code=start
func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	// a_i 尽可能小，因此用前缀数组，a_k 需要比a_i 大，因此用一个单调递减栈
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = min(prefix[i-1], nums[i])
	}

	stack := []int{}
	for i := n - 1; i > 0; i-- {
		if nums[i] > prefix[i-1] {
			for len(stack) > 0 && prefix[i-1] >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
				return true
			}
			stack = append(stack, nums[i])
		}
	}

	return false
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

