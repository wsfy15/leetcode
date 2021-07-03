/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
type node struct {
	ch   byte
	freq int
}

func frequencySort(s string) string {
	nodes := []node{}
	m := map[byte]int{}
	for _, ch := range []byte(s) {
		if _, ok := m[ch]; !ok {
			m[ch] = len(nodes)
			nodes = append(nodes, node{ch, 0})
		}

		nodes[m[ch]].freq++
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].freq > nodes[j].freq
	})

	res := []byte{}
	for _, item := range nodes {
		for i := 0; i < item.freq; i++ {
			res = append(res, item.ch)
		}
	}
	return string(res)
}

// @lc code=end

