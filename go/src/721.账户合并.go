/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
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

func merge(a, b int, uf []int) {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
}

func accountsMerge(accounts [][]string) [][]string {
	// 为每个邮箱赋一个编号
	index := 0
	email2index := map[string]int{}
	email2name := map[string]string{}
	for _, account := range accounts {
		for _, email := range account[1:] {
			if _, ok := email2index[email]; !ok {
				email2name[email] = account[0]
				email2index[email] = index
				index++
			}
		}
	}

	// 合并邮箱
	uf := NewUF(index)
	for _, account := range accounts {
		firstIndex := email2index[account[1]]
		for _, email := range account[2:] {
			merge(email2index[email], firstIndex, uf)
		}
	}

	index2email := map[int][]string{}
	for email, index := range email2index {
		parent := find(index, uf)
		index2email[parent] = append(index2email[parent], email)
	}

	res := [][]string{}
	for _, value := range index2email {
		sort.Strings(value)

		res = append(res, []string{email2name[value[0]]})
		res[len(res)-1] = append(res[len(res)-1], value...)
	}

	return res
}

// @lc code=end

