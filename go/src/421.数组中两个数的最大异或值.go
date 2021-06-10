/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
var used int
var trie [][]int

const TRIE_LEN = 1e6

func init() {
	trie = make([][]int, TRIE_LEN)
	for i := 0; i < TRIE_LEN; i++ {
		trie[i] = make([]int, 2)
	}
}

func findMaximumXOR(nums []int) int {
	used = 1
	for i := 0; i < TRIE_LEN; i++ {
		trie[i][0], trie[i][1] = 0, 0
	}

	res := 0
	for _, num := range nums {
		add(trie, num)
		res = max(res, findMax(trie, num))
	}

	return res
}

func add(trie [][]int, num int) {
	index := 0
	for i := 30; i >= 0; i-- {
		a := (num >> i) & 1
		if trie[index][a] == 0 {
			trie[index][a] = used
			used++
		}
		index = trie[index][a]
	}
}

func findMax(trie [][]int, num int) int {
	res, index := 0, 0
	for i := 30; i >= 0; i-- {
		a := (num >> i) & 1
		b := 1 - a
		if trie[index][b] != 0 {
			res |= 1 << i
			index = trie[index][b]
		} else {
			index = trie[index][a]
		}
	}

	return res
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

