/*
 * @lc app=leetcode.cn id=1331 lang=golang
 *
 * [1331] 数组序号转换
 */

// @lc code=start
func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	tmp := make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(tmp)
	m := map[int]int{}
	m[tmp[0]] = 1
	index := 2
	for i := 1; i < len(tmp); i++ {
		if tmp[i] != tmp[i-1] {
			m[tmp[i]] = index
			index++
		}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = m[arr[i]]
	}

	return arr
}

// @lc code=end

