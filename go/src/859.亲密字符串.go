/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	n, first := len(A), 0
	for first < n && A[first] == B[first] {
		first++
	}

	if first == n {
		m := map[byte]int{}
		for i := 0; i < n; i++ {
			m[A[i]]++
			if m[A[i]] > 1 {
				return true
			}
		}
		return false
	}

	second := first + 1
	for second < n && A[second] == B[second] {
		second++
	}
	if second == n {
		return false
	}
	for i := second + 1; i < n; i++ {
		if A[i] != B[i] {
			return false
		}
	}
	return A[first] == B[second] && A[second] == B[first]
}
// @lc code=end

