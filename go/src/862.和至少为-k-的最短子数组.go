/*
 * @lc app=leetcode.cn id=862 lang=golang
 *
 * [862] 和至少为 K 的最短子数组
 */

// @lc code=start
func shortestSubarray(A []int, K int) int {
	n := len(A)
	res := n + 1
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + A[i-1]
	}

	deque := []int{}
	for i := 0; i <= n; i++ {
		// prefix[x2] <= prefix[x1]，说明x1到x2这部分元素的和小于等于0，就不用考虑x1
		// 如果后面有x3 x1到x3大于等于k，x2到x3也肯定大于等于K
		for len(deque) > 0 && prefix[i] <= prefix[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		// x1到x2满足大于等于K后，x1就可以抛弃了，因为就算后面的x3满足x1到x3大于等于K，但长度更长
		for len(deque) > 0 && prefix[i]-prefix[deque[0]] >= K {
			res = min(res, i-deque[0])
			deque = deque[1:]
		}

		deque = append(deque, i)
	}

	if res == n+1 {
		return -1
	}
	return res
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

