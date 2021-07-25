/*
 * @lc app=leetcode.cn id=1743 lang=golang
 *
 * [1743] 从相邻元素对还原数组
 */

// @lc code=start
func restoreArray(adjacentPairs [][]int) []int {
	g := map[int][]int{}
	for _, pair := range adjacentPairs {
		addNode(g, pair[0], pair[1])
		addNode(g, pair[1], pair[0])
	}

	cur := getStart(g)
	last := cur
	res := []int{cur}
	for len(res) < len(adjacentPairs)+1 {
		for _, v := range g[cur] {
			if v != last {
				last = cur
				cur = v
				break
			}
		}

		res = append(res, cur)
	}

	return res
}

func addNode(g map[int][]int, start, end int) {
	for _, num := range g[start] {
		if num == end {
			return
		}
	}

	g[start] = append(g[start], end)
}

func getStart(g map[int][]int) int {
	for k, v := range g {
		if len(v) == 1 {
			return k
		}
	}
	return -1
}

// @lc code=end

