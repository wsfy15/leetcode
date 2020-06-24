/*
 * @lc app=leetcode.cn id=1431 lang=golang
 *
 * [1431] 拥有最多糖果的孩子
 */

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	num := max(candies...)
	res := make([]bool, len(candies))
	for i, candy := range candies {
		if candy + extraCandies >= num {
			res[i] = true
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

