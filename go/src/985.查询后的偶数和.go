/*
 * @lc app=leetcode.cn id=985 lang=golang
 *
 * [985] 查询后的偶数和
 */

// @lc code=start
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	cur := 0
	for _, num := range A {
		if num & 1 == 0 {
			cur += num
		}
	}

	res := make([]int, len(queries))
	for index, query := range queries {
		old := A[query[1]]
		A[query[1]] += query[0]
		if old & 1 == 0 {
			cur -= old
		}
		if A[query[1]] & 1 == 0 {
			cur += A[query[1]]
		}

		res[index] = cur
	}

	return res
}
// @lc code=end

