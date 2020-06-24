/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */

// @lc code=start
func find(uf []int, num int) int {
	for uf[num] != num {
		num = uf[num]
	}
	return num
}

func equationsPossible(equations []string) bool {
	uf := make([]int, 26) // 并查集，每个节点只记录其直接父节点编号，而不是root编号
	for i := 0; i < 26; i++ {
		uf[i] = i
	}
	for _, eq := range equations {
		if eq[1] == '=' {
			l, r := find(uf, int(eq[0] - 'a')), find(uf, int(eq[3] - 'a')) 
			if l != r {
				uf[r] = l
			}
		}
	}

	for _, eq := range equations {
		if eq[1] == '!' {
			l, r := find(uf, int(eq[0] - 'a')), find(uf, int(eq[3] - 'a')) 
			if uf[l] == uf[r] {
				return false
			}
		}
	}
	return true
}
// @lc code=end

