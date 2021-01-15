/*
 * @lc app=leetcode.cn id=1202 lang=golang
 *
 * [1202] 交换字符串中的元素
 */

// @lc code=start
func initUF(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(index int, uf []int) int {
	tmp := index
	for uf[index] != index {
		index = uf[index]
	}
	for uf[tmp] != index {
		tmp, uf[tmp] = uf[tmp], index
	}
	return index
}

func merge(a, b int, uf []int) bool {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
	return aroot != broot
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	// 并查集得到每个字符所属的组别，然后对每个组进行排序，再根据每个位置所属组别构造结果
	n := len(s)
	part := n
	uf := initUF(n)
	for _, pair := range pairs {
		if merge(pair[0], pair[1], uf) {
			part--
		}
	}

	for i := 0; i < n; i++ { // 保证每个节点都指向其root
		find(i, uf)
	}

	bytes := make([][]byte, part)
	used := make([]int, part) // 每个组当前可以使用的字符下标
	uf2index := map[int]int{} // 每个uf 的组别
	index := 0                // 组别
	for i, ch := range []byte(s) {
		if _, ok := uf2index[uf[i]]; !ok {
			uf2index[uf[i]] = index
			index++
		}

		bytes[uf2index[uf[i]]] = append(bytes[uf2index[uf[i]]], ch)
	}

	for i := range bytes {
		sort.Slice(bytes[i], func(a, b int) bool {
			return bytes[i][a] < bytes[i][b]
		})
	}

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		group := uf2index[uf[i]]
		res[i] = bytes[group][used[group]]
		used[group]++
	}
	return string(res)
}

// @lc code=end

