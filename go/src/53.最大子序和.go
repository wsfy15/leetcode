/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	max, sum := nums[0], 0
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}

		if max < sum {
			max = sum
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	res, max := nums[0], nums[0]
	size := len(nums)
	for i := 1; i < size; i++ {
		if res < 0 {
			if nums[i] > res {
				res = nums[i]
			}
		} else {
			res += nums[i]
			if res < 0 {
				res = nums[i]
			}
		}

		if res > max {
			max = res
		}
	}

	return max
}
// @lc code=end

