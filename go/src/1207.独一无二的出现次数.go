/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	counter := map[int]int{}
	unique := map[int]bool{}
	for _, num := range arr {
		counter[num]++
	}
	for _, v := range counter {
		if _, ok := unique[v]; ok {
			return false
		}
		unique[v] = true
	}
	return true
}

// @lc code=end

