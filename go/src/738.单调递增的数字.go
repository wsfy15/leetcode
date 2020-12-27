/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}

	nums := []int{}
	tmp := N
	for tmp > 0 {
		num := tmp % 10
		nums = append(nums, num)
		tmp /= 10
	}

	// 从高位往地位遍历，遇到不满足递增的就结束，否则记录最大值（首次出现）的位置
	inOrder := true
	maxIndex := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			maxIndex = i
		} else if nums[i] < nums[i+1] {
			inOrder = false
			break
		}
	}

	if inOrder {
		return N
	}

	nums[maxIndex]--
	for i := 0; i < maxIndex; i++ {
		nums[i] = 9
	}

	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		res *= 10
		res += nums[i]
	}

	return res
}

// @lc code=end

