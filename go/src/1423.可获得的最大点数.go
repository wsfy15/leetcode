/*
 * @lc app=leetcode.cn id=1423 lang=golang
 *
 * [1423] 可获得的最大点数
 */

// @lc code=start
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	nums := make([]int, k + k)
	
	for i := 0; i < k; i++ {
		nums[i] = cardPoints[n - k + i]
	}
	for i := k; i < k + k; i++ {
		nums[i] = cardPoints[i - k]
	}

	cur, res := 0, 0
	for i := 0; i < k - 1; i++ {
		cur += nums[i]
	} 
	for i := k - 1; i < k + k; i++ {
		cur += nums[i]
		res = max(res, cur)
		cur -= nums[i - k + 1]
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

