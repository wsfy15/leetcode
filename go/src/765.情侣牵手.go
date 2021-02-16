/*
 * @lc app=leetcode.cn id=765 lang=golang
 *
 * [765] 情侣牵手
 */

// @lc code=start
func NewUF(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(n int, uf []int) int {
	if n != uf[n] {
		uf[n] = find(uf[n], uf)
	}
	return uf[n]
}

func merge(a, b int, uf []int) bool {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
	return aroot != broot
}

// 人的编号/2 作为每对情侣的编号
func minSwapsCouples(row []int) int {
	n := len(row)
	m := n / 2
	uf := NewUF(m)

	for i := 0; i < n; i += 2 {
		if merge(row[i] / 2, row[i + 1] / 2, uf) {
			m--
		}
	}

	return n / 2 - m
}
// @lc code=end

