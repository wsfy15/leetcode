/*
 * @lc app=leetcode.cn id=1744 lang=golang
 *
 * [1744] 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
 */

// @lc code=start
func canEat(candiesCount []int, queries [][]int) []bool {
	n, m := len(queries), len(candiesCount)
	res := make([]bool, n)

	prefix := make([]int, m+1)
	for i := 1; i <= m; i++ {
		prefix[i] = prefix[i-1] + candiesCount[i-1]
	}

	for i, q := range queries {
		typ, day := q[0], q[1]+1
		// 每天吃1颗，不会在day天前把typ糖果都吃掉
		// 每天吃cap颗，在第day天可以达到第typ种糖果
		res[i] = day <= prefix[typ+1] && prefix[typ] < day*q[2]
	}

	return res
}

// @lc code=end

