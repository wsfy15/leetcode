/*
 * @lc app=leetcode.cn id=1707 lang=golang
 *
 * [1707] 与数组中元素的最大异或值
 */

// @lc code=start
type trie struct {
	children []*trie
	min      int
}

func NewTrie() *trie {
	return &trie{
		children: make([]*trie, 2),
		min:      math.MaxInt32,
	}
}

func (t *trie) add(num int) {
	cur := t
	cur.min = min(cur.min, num)
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		if cur.children[bit] == nil {
			cur.children[bit] = NewTrie()
		}

		cur = cur.children[bit]
		cur.min = min(cur.min, num)
	}
}

func (t *trie) getMaxXor(val, limit int) int {
	cur := t
	if cur.min > limit {
		return -1
	}

	res := 0
	for i := 31; i >= 0; i-- {
		bit := (val >> i) & 1
		if cur.children[bit^1] != nil && cur.children[bit^1].min <= limit {
			res |= 1 << i
			bit ^= 1
		}

		cur = cur.children[bit]
	}

	return res
}

func maximizeXor(nums []int, queries [][]int) []int {
	n := len(queries)
	res := make([]int, n)
	t := NewTrie()
	for _, num := range nums {
		t.add(num)
	}

	for i, q := range queries {
		res[i] = t.getMaxXor(q[0], q[1])
	}
	return res
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

