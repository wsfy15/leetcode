/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
// 遇到递减时，有两种情况：1、3、2  和 5、6、4，第一种可以将3改为2，第二种将3改为6
func checkPossibility(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return true
	}

	count := 0
	if nums[0] > nums[1] {
		nums[0] = nums[1]
		count = 1
	}

	for i := 1; i < n - 1; i++ {
		if nums[i] > nums[i + 1] {
			count++
			if count > 1 {
				return false
			}

			left := nums[i - 1]
			if nums[i + 1] > left {
				nums[i] = nums[i + 1]
			} else {
				nums[i + 1] = nums[i]
			}
		}
	}

	return true
}
// @lc code=end

