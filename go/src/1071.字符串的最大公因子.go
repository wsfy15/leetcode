/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] 字符串的最大公因子
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	// 如果存在最大公因子，则str1和str2都是该公因子的重复，str1+str2也是公因子的重复
	if str1 + str2 != str2 + str1 { 
		return ""
	}

	// 既然存在gcd，那么gcd子串的长度肯定是str1和str2的最大公约数
	len1 := len(str1)
	len2 := len(str2)

	num := gcd(len1, len2)
	return str1[:num]
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for x % y != 0 {
		x, y = y, x % y
	}

	return y
}
// @lc code=end

