/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}

	if dividend == math.MinInt32 {
		if divisor < 0 {
			return 1 + divide(-(dividend - divisor), -divisor)
		} else {
			return -1 - divide(-(dividend + divisor), divisor)
		}
	}

	if divisor == math.MinInt32 {
		if dividend != divisor {
			return 0
		}
		return 1
	}

	if dividend < 0 && divisor < 0 {
		return divide(-dividend, -divisor)
	}

	if dividend < 0 {
		return -divide(-dividend, divisor)
	}

	if divisor < 0 {
		return -divide(dividend, -divisor)
	}

	return int(div(int32(dividend), int32(divisor)))
}

func div(dividend, divisor int32) int32 {
	if dividend < divisor {
		return 0
	}

	var cur, k int32 = divisor, 1
	var count int32 = 0
	for dividend > 0 {
		count += k
		dividend -= cur
		if dividend < cur || dividend < cur << 1 { // 防止溢出
			return count + div(dividend, divisor)
		}
		k <<= 1
		cur <<= 1
	}

	if dividend < 0 {
		k >>= 1
		cur >>= 1
		count -= k
		dividend += cur
	}

	return count + div(dividend, divisor)
}

func divide2(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		} 
		return math.MaxInt32
	}

	if dividend < 0 && divisor < 0 {
		return divide(-dividend, -divisor)
	}

	if dividend < 0 {
		return -divide(-dividend, divisor)
	}

	if divisor < 0 {
		return -divide(dividend, -divisor)
	}

	count, cur := 0, 0
	for cur <= dividend {
		count++
		cur += divisor 
	}

	return count - 1
}
// @lc code=end

