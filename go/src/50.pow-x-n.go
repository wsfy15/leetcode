/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}

	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	res := 1.0
	i := x // i 表示x的幂
	// 对n的二进制位从低位往高位处理
	for n > 0 {
		if n & 1 == 1 {
			res *= i
		}
		i *= i // x^1 x^2 x^4 x^8
		n >>= 1
	}

	return res
}

func myPow2(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}

	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	tmp := myPow(x, n / 2)
	res := tmp * tmp
	if n & 1 == 1 { // 奇数
		res *= x
	}
	return res
}



// @lc code=end

