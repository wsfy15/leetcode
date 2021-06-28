/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */

// @lc code=start
func maxPoints(points [][]int) int {
	n, res := len(points), 0
	if n < 3 {
		return n
	}

	// 遍历每个点，计算该点与其它点组合的斜率，计算有多少点都在同一直线上
	for i := 0; i < n; i++ {
		// 以 分子@分母 为key，	value为点的数量
		m := map[string]int{}
		duplicate, cur := 0, 0
		for j := i + 1; j < n; j++ {
			y, x := points[j][1]-points[i][1], points[j][0]-points[i][0]
			if x == 0 && y == 0 {
				duplicate++
				continue
			}

			g := gcd(y, x)
			x, y = x/g, y/g
			key := strconv.Itoa(x) + "@" + strconv.Itoa(y)
			m[key]++
			cur = max(cur, m[key])
		}

		res = max(res, cur+duplicate+1)
	}

	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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

