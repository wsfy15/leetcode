/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 */

// @lc code=start
func reorganizeString(S string) string {
	counter := make([]int, 26)
	for _, ch := range S {
		counter[ch-'a']++
	}
	if max(counter...) > (len(S)+1)/2 {
		return ""
	}

	res := make([]byte, len(S))
	odd, even := 1, 0
	for i := 0; i < 26; i++ {
		for counter[i] > 0 && counter[i] < len(S)/2+1 && odd < len(S) {
			res[odd] = 'a' + byte(i)
			counter[i]--
			odd += 2
		}

		for counter[i] > 0 {
			res[even] = 'a' + byte(i)
			counter[i]--
			even += 2
		}
	}
	return string(res)
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

