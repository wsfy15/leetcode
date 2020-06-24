/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	x := uint32(num)
	bytes := []byte{}
	for x != 0 {
		var cur byte = 0
		for j := 0; j < 4; j++ {
			if x & (1 << j) > 0 {
				cur += 1 << j
			}
		}
		if cur > 9 {
			cur = 'a' + cur - 10
		} else {
			cur = '0' + cur
		}
		bytes = append(bytes, cur)
		x >>= 4
	}

	i, j := 0, len(bytes) - 1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
// @lc code=end

