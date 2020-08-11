/*
 * @lc app=leetcode.cn id=1394 lang=golang
 *
 * [1394] 找出数组中的幸运数
 */

// @lc code=start
func findLucky(arr []int) int {
	counter := map[int]int{}
	for _, num := range arr {
		counter[num]++
	}

	res := -1
	for k, v := range counter {
		if k == v && k > res {
			res = k
		}
	}
	return res
}

// @lc code=end

