/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if v == first || v == second || v == third {
			continue
		}

		if v > first {
			first, second, third = v, first, second
		} else if v > second {
			second, third = v, second
		} else if v > third {
			third = v
		}
	}
	
	if third == math.MinInt64 {
		return first
	}
	return third
}
// @lc code=end

