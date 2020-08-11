/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	counter := [26]int{}
	for i := 0; i < len(text); i++ {
		counter[text[i]-'a']++
	}

	res := min(counter['a'-'a'], counter['b'-'a'], counter['n'-'a'], counter['l'-'a']/2, counter['o'-'a']/2)
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

