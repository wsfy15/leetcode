/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 */

// @lc code=start
func commonChars(A []string) []string {
	count := make([]int, 26)
	for i := 0; i < len(A[0]); i++ {
		count[A[0][i] - 'a']++
	}

	for i := 1; i < len(A); i++ {
		tmp := make([]int, 26)
		for j := 0; j < len(A[i]); j++ {
			tmp[A[i][j] - 'a']++
		}

		for j := 0; j < 26; j++ {
			count[j] = min(count[j], tmp[j])
		}
	}

	res := []string{}
	for i := 0; i < 26; i++ {
		for j := 0; j < count[i]; j++ {
			res = append(res, string('a' + i))
		}
	}

	return res
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

