/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	nums := make([]byte, m + n) // m位数与n位数相乘，结果最大位数为m+n
	// nums[0]存储最高位，nums[m+n-1]最低位
	
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 第i位与第j位相乘结果在第i+j位和第i+j+1位
			var res byte = (num1[i] - '0') * (num2[j] - '0')
			if res > 10 {
				nums[i + j] += res / 10
				res %= 10
			}
			nums[i + j + 1] += res
			pos := i + j + 1
			if nums[pos] < 10 {
				pos--
			}
			for pos >= 0 && nums[pos] >= 10 {
				nums[pos - 1] += nums[pos] / 10
				nums[pos] %= 10
				pos--
			}
		}
	}

	i := 0
	for i < m + n && nums[i] == 0 {
		i++
	}
	if i == m + n {
		return "0"
	}

	sb := strings.Builder{}
	for i < m + n {
		sb.WriteByte(nums[i] + '0')
		i++
	}
	return sb.String()
}
// @lc code=end

