/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	n := len(ratings)
	if n < 2 {
		return n
	}

	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i - 1] {
			left[i] = left[i - 1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i + 1] {
			right[i] = right[i + 1] + 1
		}
	}

	res := n
	for i := 0; i < n; i++ {
		res += max(left[i], right[i])
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

