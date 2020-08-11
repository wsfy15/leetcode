/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	n := len(num)
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if dfs(num, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func dfs(num string, i, j, k int) bool {
	if (num[i] == '0' && j - i > 1) || (num[j] == '0' && k - j > 1) {
		return false
	}

	a, b := num[i:j], num[j:k]
	c := add(a, b)
	if k + len(c) > len(num) || num[k:k+len(c)] != c {
		return false
	}

	if k + len(c) == len(num) {
		return true
	}

	return dfs(num, j, k, k + len(c))
}

func add(a, b string) string {
	n, m := len(a) - 1, len(b) - 1
	res := []byte{}
	var carry byte = 0
	for n >= 0 && m >= 0 {
		tmp := a[n] - '0' + b[m] - '0' + carry
		carry = tmp / 10
		res = append(res, '0' + tmp % 10)
		n--
		m--
	}

	for n >= 0 {
		tmp := a[n] - '0' + carry
		carry = tmp / 10
		res = append(res, '0' + tmp % 10)
		n--
	}

	for m >= 0 {
		tmp := b[m] - '0' + carry
		carry = tmp / 10
		res = append(res, '0' + tmp % 10)
		m--
	}

	if carry > 0 {
		res = append(res, '1')
	}
	
	i, j := 0, len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
// @lc code=end

