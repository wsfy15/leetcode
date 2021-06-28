/*
 * @lc app=leetcode.cn id=1239 lang=golang
 *
 * [1239] 串联字符串的最大长度
 */

// @lc code=start
func maxLength(arr []string) int {
	return dfs(arr, 0, 0)
}

func dfs(arr []string, used, index int) int {
	if index == len(arr) {
		return 0
	}

	newUsed := merge(used, arr[index])
	if newUsed > 0 {
		return max(len(arr[index])+dfs(arr, newUsed, index+1), dfs(arr, used, index+1))
	}
	return dfs(arr, used, index+1)
}

func merge(used int, str string) int {
	res := used
	for _, ch := range str {
		// 已经存在
		if res&(1<<(ch-'a')) > 0 {
			return -1
		}
		res |= 1 << (ch - 'a')
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

