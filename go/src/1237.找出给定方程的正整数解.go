/*
 * @lc app=leetcode.cn id=1237 lang=golang
 *
 * [1237] 找出给定方程的正整数解
 */

// @lc code=start
/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	// 因为f(x, y) > 0，所以x、y的上界设为z即可
	x, y := 1, 1000
	res := [][]int{}
	for x <= 1000 && y >= 1 {
		tmp := customFunction(x, y)
		if tmp == z {
			res = append(res, []int{x, y})
			x--
			y--
		} else if tmp < z {
			x++
		} else {
			y--
		}
	}

	return res
}

// @lc code=end

