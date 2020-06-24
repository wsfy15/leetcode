/*
 * @lc app=leetcode.cn id=686 lang=golang
 *
 * [686] 重复叠加字符串匹配
 */

// @lc code=start
// 生成的新字符串 最大长度为 LEN(B) + 2 * LEN(A), 即考虑A的后部分作为B的开头，A的前部分作为B的结尾，剩下全是重复的完整的A
// A`(AAAA...)A`
func repeatedStringMatch(A string, B string) int {
	n, m := len(A), len(B)
	newA := A
	maxLen := 2 * n + m
	for count := 1; len(newA) <= maxLen; count++ {
		if len(newA) >= len(B) && strings.Contains(newA, B) {
			return count
		}
		newA += A
	} 

	return -1
}
// @lc code=end

