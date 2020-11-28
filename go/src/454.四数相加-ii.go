/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	ab := map[int]int{}
	for _, a := range A {
		for _, b := range B {
			ab[a+b]++
		}
	}

	res := 0
	for _, c := range C {
		for _, d := range D {
			if v, ok := ab[-(c + d)]; ok {
				res += v
			}
		}
	}

	return res
}

// @lc code=end

