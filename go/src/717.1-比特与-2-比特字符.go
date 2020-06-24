/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1比特与2比特字符
 */

// @lc code=start
func isOneBitCharacter2(bits []int) bool {
	n, i := len(bits), 0
	for ; i < n - 1; i++ {
		if bits[i] == 1 {
			i++
		}
	}
	return i == n - 1 && bits[n - 1] == 0
}

func isOneBitCharacter(bits []int) bool {
	// 只要考虑最后的0前面的连续1的个数
	n, count := len(bits), 0
	if n == 0 && bits[n - 1] != 0 {
		return false
	}

	for i := n - 2; i >= 0 && bits[i] == 1; i-- {
		count++
	}
	return count & 1 == 0
}
// @lc code=end

