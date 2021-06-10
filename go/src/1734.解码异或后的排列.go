/*
 * @lc app=leetcode.cn id=1734 lang=golang
 *
 * [1734] 解码异或后的排列
 */

// @lc code=start
func decode(encoded []int) []int {
	//  perm 是前 n 个正整数的排列 perm=[a,b,c,d,e] encoded=[f,g,h,i]
	// 可以求出 前n个正整数的异或  a^b^c^d^e
	//  f=a^b g=b^c h=c^d i=d^e
	// g ^ i = b^c^d^e
	// 由此可以得到第一个数a

	n := len(encoded) + 1
	a := 0
	for i := 1; i <= n; i++ {
		a ^= i
	}

	for i := n - 2; i > 0; i -= 2 {
		a ^= encoded[i]
	}

	res := make([]int, n)
	res[0] = a
	for i := 1; i < n; i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}

	return res
}

// @lc code=end

