/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	nums := make([]int, n)
	nums[0] = 1
	pointers := make([]int, len(primes))
	for i := 1; i < n; i++ {
		tmp := make([]int, len(primes))
		for i, v := range primes {
			tmp[i] = v * nums[pointers[i]]
		}

		target := min(tmp...)
		nums[i] = target
		for i, v := range tmp {
			if v == target {
				pointers[i]++
			}
		}
	}

	return nums[n - 1]
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

