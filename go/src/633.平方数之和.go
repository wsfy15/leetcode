/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		num := i * i + j * j
		if num == c {
			return true
		} else if num < c {
			i++
		} else {
			j--
		}
	}
	return false
}

func judgeSquareSum2(c int) bool {
	counter := map[int]bool{}
	counter[0] = true
	for i := 1; ; i++ {
		num := i * i
		counter[num] = true

		if num > c {
			return false
		}

		if _, ok := counter[c - num]; ok {
			return true
		}
	}
}
// @lc code=end

