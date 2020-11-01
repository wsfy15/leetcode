/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(S string) []int {
	n := len(S)
	index := make([]int, 26) // 记录每个字母最后出现位置
	for i, ch := range S {
		index[ch-'a'] = i
	}

	res := []int{}
	start, end, cur := 0, 0, 0
	for cur < n {
		end = max(end, index[S[cur]-'a'])
		if end == cur {
			res = append(res, end-start+1)
			start = cur + 1
		}
		cur++
	}
	return res
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

