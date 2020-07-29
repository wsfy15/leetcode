/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 */

// @lc code=start
func addToArrayForm(A []int, K int) []int {
	n, carry := len(A), 0
	for i := n - 1; i >= 0; i-- {
		num := K % 10 + A[i] + carry
		carry = 0
		if num > 9 {
			carry = 1
			num %= 10
		}

		A[i] = num
		K /= 10
	}

	if carry > 0 || K > 0 {
		K += carry
		tmp := []int{}
		for K > 0 {
			tmp = append(tmp, K % 10)
			K /= 10
		}
		i, j := 0, len(tmp) - 1
		for i < j {
			tmp[i], tmp[j] = tmp[j], tmp[i]
			i++
			j--
		}

		A = append(tmp, A...)
	}

	return A
}
// @lc code=end

