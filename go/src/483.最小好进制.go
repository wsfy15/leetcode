/*
 * @lc app=leetcode.cn id=483 lang=golang
 *
 * [483] 最小好进制
 */

// @lc code=start
func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	res := num
	// 最大的1e18，以二进制表示，1的个数不会超过64
	// 当i比较大时，getNum会溢出
	for i := 1; i <= 64; i++ {
		l, r := 2, num // x进制
		for l < r {
			mid := l + (r-l)>>1
			tmp := getNum(mid, i)
			if tmp == num {
				res = min(res, mid)
				break
			} else if tmp < num {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}

	return strconv.Itoa(res)
}

// n位x进制，每一位都是1
func getNum(x, n int) int {
	num := 0
	for i := 0; i < n; i++ {
		num = num*x + 1
	}
	return num
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

