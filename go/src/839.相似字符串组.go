/*
 * @lc app=leetcode.cn id=839 lang=golang
 *
 * [839] 相似字符串组
 */

// @lc code=start
func NewUF(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(node int, uf []int) int {
	if uf[node] != node {
		uf[node] = find(uf[node], uf)
	}
	return uf[node]
}

func merge(a, b int, uf []int) bool {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
	return aroot != broot
}

func isSimilar(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return true
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := NewUF(n)
	count := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) && merge(i, j, uf) {
				count--
			}
		}
	}

	return count
}
// @lc code=end

