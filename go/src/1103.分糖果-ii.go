/*
 * @lc app=leetcode.cn id=1103 lang=golang
 *
 * [1103] 分糖果 II
 */

// @lc code=start
func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	cur, i := 1, 0
	for candies > 0 {
		if candies >= cur {
			res[i] += cur
			candies -= cur
			i = (i + 1) % num_people
			cur++
		} else {
			res[i] += candies
			return res
		}
	}
	return res
}
// @lc code=end

