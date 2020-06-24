/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	recorder, exist := map[int]bool{}, map[string]bool{}
	n, cur := len(s), 0
	if n < 10 {
		return nil
	}

	m := map[byte]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}

	for i := 0; i < 10; i++ {
		cur <<= 2
		cur |= m[s[i]]
	}
	recorder[cur] = true

	for i := 10; i < n; i++ {
		cur <<= 2
		cur |= m[s[i]]
		cur &= 1 << 20 - 1
		if _, ok := recorder[cur]; ok {
			exist[s[i - 9: i + 1]] = true
		}
		recorder[cur] = true
	}

	res := []string{}
	for k, _ := range exist {
		res = append(res, k)
	}
	return res
}
// @lc code=end

