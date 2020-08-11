/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
type pair struct {
	index int
	num   int
}

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	pairs := make([]pair, n)
	for index, num := range nums {
		pairs[index] = pair{index: index, num: num}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].num < pairs[j].num
	})

	for i := 0; i < n; i++ {
		res[pairs[i].index] = i
		j := i + 1
		for j < n && pairs[j].num == pairs[i].num {
			res[pairs[j].index] = i
			j++
		}
		i = j - 1
	}

	return res
}

// @lc code=end

