/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	
	// 假设当前要求爬到三楼的步数，定义x为爬到一楼（3 - 2）的步数 y：爬到二楼（3 - 1）的步数
	x, y := 1, 2 
	for i := 3; i <= n; i++ {
		x, y = y, x + y // x更新为爬到其下一层的步数（即y），y更新为从x和y到 y+1 (x+2)层的步数
	}
	return y
}


func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	var count = make([]int, n + 1)
	count[1] = 1
	count[2] = 2

	return util(n, count)
}

func util(n int , count []int) int {
	if count[n] != 0 {
		return count[n]
	}
	
	count[n] = util(n - 1, count) + util(n - 2, count)
	return count[n] 
}
// @lc code=end

