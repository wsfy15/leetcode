/*
 * @lc app=leetcode.cn id=947 lang=golang
 *
 * [947] 移除最多的同行或同列石头
 */

// @lc code=start
func NewUf(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(root int, uf []int) int {
	tmp := root
	for root != uf[root] {
		root = uf[root]
	}

	for tmp != uf[tmp] {
		tmp, uf[tmp] = uf[tmp], root
	}
	return root
}

func merge(a, b int, uf []int) {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
}

func removeStones(stones [][]int) int {
	n := len(stones)
	row, col := make([][]int, 10001), make([][]int, 10001)
	for i, stone := range stones {
		row[stone[0]] = append(row[stone[0]], i)
		col[stone[1]] = append(col[stone[1]], i)
	}

	uf := NewUf(n)
	for i := 0; i <= 10000; i++ {
		if len(row[i]) == 0 {
			continue
		}
		root := row[i][0]
		for _, index := range row[i][1:] {
			merge(index, root, uf)
		}
	}

	for i := 0; i <= 10000; i++ {
		if len(col[i]) == 0 {
			continue
		}
		root := col[i][0]
		for _, index := range col[i][1:] {
			merge(index, root, uf)
		}
	}

	counter := map[int]struct{}{}
	for i := 0; i < n; i++ {
		counter[find(i, uf)] = struct{}{}
	}

	return n - len(counter)
}

// @lc code=end

