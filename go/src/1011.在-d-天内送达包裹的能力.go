/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
func shipWithinDays(weights []int, D int) int {
	sum, n := 0, len(weights)
	for i := 0; i < n; i++ {
		sum += weights[i]
	}

	l, r := sum/D, sum
	for l < r {
		mid := l + (r-l)>>1
		if valid(weights, D, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func valid(weights []int, D, W int) bool {
	cur, count := 0, 1
	for _, weight := range weights {
		if weight > W {
			return false
		}
		if cur+weight > W {
			count++
			cur = weight
		} else {
			cur += weight
		}
	}

	return count <= D
}

// @lc code=end

