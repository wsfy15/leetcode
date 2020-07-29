/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
// 只要关注个位上的数是否为0或5就行
func prefixesDivBy5(A []int) []bool {
	n := len(A)
	res := make([]bool, n)
	num := 0
	for i, v := range A {
		num = ((num << 1) + v) % 10
		res[i] = (num % 5) == 0
	}
	return res
}
// @lc code=end

