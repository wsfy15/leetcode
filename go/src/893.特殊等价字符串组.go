/*
 * @lc app=leetcode.cn id=893 lang=golang
 *
 * [893] 特殊等价字符串组
 */

// @lc code=start
func numSpecialEquivGroups(A []string) int {
	n := len(A)
	trans := make([]string, n)
	for i := 0; i < n; i++ {
		trans[i] = process(A[i])
	}

	counter := map[string]bool{}
	for i := 0; i < n; i++ {
		counter[trans[i]] = true
	}
	return len(counter)
}

// 分别把偶数和奇数位置的字符排序后放在前半部分和后半部分
func process(str string) string {
	n := len(str)
	bytes, j := make([]byte, n), 0
	for i := 0; i < n; i += 2 {
		bytes[j] = str[i]
		j++
	}
	sort.Slice(bytes[:j], func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	tmp := j
	for i := 1; i < n; i += 2 {
		bytes[j] = str[i]
		j++
	}
	sort.Slice(bytes[tmp:], func(i, j int) bool {
		return bytes[i + tmp] < bytes[j + tmp]
	})
	return string(bytes)
}
// @lc code=end

