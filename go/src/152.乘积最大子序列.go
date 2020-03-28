	/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */

// @lc code=start
import "math"

// 因为存在负负得正的情况，所以需要记录最大值和最小值（负数）
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	curMax := math.MinInt32
	imax := 1 // 以当前元素为结尾的子序列的 最大乘积
	imin := 1 // 以当前元素为结尾的子序列的 最小乘积
	// 当遇到负数时，交换imax和imin
	for i := range nums {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = max(imax * nums[i], nums[i])
		imin = min(imin * nums[i], nums[i])
		curMax = max(imax, curMax)
	}

	return curMax
}

func max(i ,j int) int {
	if i > j{
		return i
	}
	return j
}

func min(i ,j int) int {
	if i > j{
		return j
	}
	return i
}
// @lc code=end

