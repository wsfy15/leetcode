/*
 * @lc app=leetcode.cn id=1337 lang=golang
 *
 * [1337] 方阵中战斗力最弱的 K 行
 */

// @lc code=start
type pair struct {
	ones int
	line int
}

func kWeakestRows(mat [][]int, k int) []int {
	res := make([]int, k)
	pairs := make([]pair, len(mat))
	for index, row := range mat {
		num := 0
		for _, v := range row {
			if v == 1 {
				num++
			} else {
				break
			}
		}
		pairs[index] = pair{ones: num, line: index}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].ones == pairs[j].ones {
			return pairs[i].line < pairs[j].line
		}
		return pairs[i].ones < pairs[j].ones
	})

	for i := 0; i < k; i++ {
		res[i] = pairs[i].line
	}
	return res
}

// @lc code=end

