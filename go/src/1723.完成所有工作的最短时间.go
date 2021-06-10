/*
 * @lc app=leetcode.cn id=1723 lang=golang
 *
 * [1723] 完成所有工作的最短时间
 */

// @lc code=start
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	tot := make([]int, 1<<n) // i的二进制位表示一个jobs子集，tot[i] 表示完成该子集需要的时间
	for i := 1; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if i&(1<<j) == 0 {
				continue
			}

			tot[i] = tot[i-(1<<j)] + jobs[j]
			break
		}
	}

	sum := 0
	for _, job := range jobs {
		sum += job
	}

	l, r := max(jobs...), sum
	for l < r {
		mid := l + (r-l)/2

		// dp[i] 表示当最大工作时间为mid时，完成i子集所需工人数量
		dp := make([]int, 1<<n)
		for i := 1; i < 1<<n; i++ {
			dp[i] = r
		}

		for i := 1; i < 1<<n; i++ {
			for cur := i; cur > 0; cur = (cur - 1) & i {
				if tot[cur] <= mid {
					dp[i] = min(dp[i], 1+dp[i-cur])
				}
			}
		}

		if dp[1<<n-1] <= k {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func getWorkers(num int, jobs []int) int {
	cur, res := 0, 1
	for _, job := range jobs {
		if cur+job > num {
			cur = 0
			res++
		}
		cur += job
	}
	return res
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
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

