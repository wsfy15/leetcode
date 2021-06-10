/*
 * @lc app=leetcode.cn id=1482 lang=golang
 *
 * [1482] 制作 m 束花所需的最少天数
 */

// @lc code=start
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}

	l, r := 0, max(bloomDay...)
	for l < r {
		mid := l + (r-l)>>1

		count := getCount(bloomDay, mid, k)
		if count >= m {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// 使用相邻的k朵花，最多days天，可以制作多少束花
func getCount(bloomDay []int, days int, k int) int {
	cur, res := 0, 0
	for _, num := range bloomDay {
		if num <= days {
			cur++
			if cur == k {
				cur = 0
				res++
			}
		} else {
			cur = 0
		}
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

// @lc code=end

