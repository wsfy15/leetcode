/*
 * @lc app=leetcode.cn id=916 lang=golang
 *
 * [916] 单词子集
 */

// @lc code=start
func wordSubsets(A []string, B []string) []string {
	// 计数 B 中每个字母出现的最大次数
	counter := make([]int, 26)

	for _, word := range B {
		cur := make([]int, 26)
		for _, ch := range word {
			cur[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			counter[i] = max(counter[i], cur[i])
		}
	}

	res := []string{}
	for _, word := range A {
		cur := make([]int, 26)
		flag := true
		for _, ch := range word {
			cur[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if counter[i] > cur[i] {
				flag = false
				break
			}
		}

		if flag {
			res = append(res, word)
		}
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

