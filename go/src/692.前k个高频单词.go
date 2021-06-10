/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
type node struct {
	word string
	freq int
}

func isBigger(a, b node) bool {
	return a.freq > b.freq || (a.freq == b.freq && a.word < b.word)
}

type heap struct {
	nodes []node
	count int
	index map[string]int
}

func NewHeap() *heap {
	return &heap{
		nodes: make([]node, 1),
		index: make(map[string]int),
	}
}

func (h *heap) push(word string) {
	if _, ok := h.index[word]; !ok {
		h.nodes = append(h.nodes, node{word, 0})
		h.count++
		h.index[word] = h.count
	}

	index := h.index[word]
	h.nodes[index].freq++
	for index/2 > 0 && isBigger(h.nodes[index], h.nodes[index/2]) {
		h.nodes[index], h.nodes[index/2] = h.nodes[index/2], h.nodes[index]
		h.index[h.nodes[index].word] = index
		h.index[h.nodes[index/2].word] = index / 2
		index /= 2
	}
}

func (h *heap) pop() string {
	res := h.nodes[1].word
	delete(h.index, res)
	h.nodes[1] = h.nodes[h.count]
	h.index[h.nodes[1].word] = 1
	h.count--

	index := 1
	for {
		pos := index
		if index*2 <= h.count && isBigger(h.nodes[index*2], h.nodes[index]) {
			pos = index * 2
		}
		if index*2+1 <= h.count && isBigger(h.nodes[index*2+1], h.nodes[pos]) {
			pos = index*2 + 1
		}

		if pos == index {
			break
		}

		h.nodes[index], h.nodes[pos] = h.nodes[pos], h.nodes[index]
		// index 可以不维护
		h.index[h.nodes[index].word] = index
		h.index[h.nodes[pos].word] = pos
		index = pos
	}

	return res
}

func topKFrequent(words []string, k int) []string {
	h := NewHeap()
	for _, word := range words {
		h.push(word)
	}

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = h.pop()
	}
	return res
}

// @lc code=end

