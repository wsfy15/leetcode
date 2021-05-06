/*
 * @lc app=leetcode.cn id=972 lang=golang
 *
 * [972] 相等的有理数
 */

// @lc code=start
func isRationalEqual(s string, t string) bool {
	n1, m1 := fraction(s)
	n2, m2 := fraction(t)
	return n1*m2 == n2*m1
}

// 返回分子和分母
func fraction(s string) (int, int) {
	dotIndex := strings.Index(s, ".")
	if dotIndex == -1 || dotIndex == len(s)-1 { // 无小数部分
		if dotIndex == len(s)-1 {
			s = s[:dotIndex]
		}
		num, _ := strconv.Atoi(s)
		return num, 1
	}

	num, _ := strconv.Atoi(s[:dotIndex])
	s = s[dotIndex+1:]
	bracketIndex := strings.Index(s, "(")
	if bracketIndex == -1 { // 小数部分无循环
		decimal, _ := strconv.Atoi(s)
		if decimal == 0 { // 1.0, 1
			return num, 1
		}

		if num == 0 {
			num, decimal = decimal, int(math.Pow10(len(s)))
		} else {
			num *= int(math.Pow10(len(s)))
		}

		divisor := gcd(num, decimal)
		return num / divisor, decimal / divisor
	}

	var noCycleNum, noCyclePow int = 0, 1
	if bracketIndex > 0 {
		noCycleStr := s[:bracketIndex]
		noCycleNum, _ = strconv.Atoi(noCycleStr)
		noCyclePow = int(math.Pow10(bracketIndex))
	}

	cycleStr := s[bracketIndex+1 : len(s)-1]
	cyclePow := int(math.Pow10(len(cycleStr))) - 1
	cycleNum, _ := strconv.Atoi(cycleStr)

	num *= noCyclePow * cyclePow
	num += noCycleNum*cyclePow + cycleNum
	divisor := gcd(num, noCyclePow*cyclePow)
	return num / divisor, noCyclePow * cyclePow / divisor
}

func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}

	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// @lc code=end

