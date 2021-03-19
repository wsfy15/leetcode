/*
 * @lc app=leetcode.cn id=1052 lang=golang
 *
 * [1052] 爱生气的书店老板
 */

// @lc code=start
func maxSatisfied(customers []int, grumpy []int, X int) int {
	n, left, right, count, res := len(customers), 0, 0, 0, 0
	for right < n {
		for left + X <= right {
			if grumpy[left] == 1 {
				count -= customers[left]
			}
			left++
		}

		if grumpy[right] == 1 {
			count += customers[right]
			res = max(res, count)
		}
		right++
	}

	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			res += customers[i]
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

