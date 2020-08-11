/*
 * @lc app=leetcode.cn id=1184 lang=golang
 *
 * [1184] 公交站间的距离
 */

// @lc code=start
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}

	total := 0
	for _, num := range distance {
		total += num
	}

	res := 0
	for i := start; i < destination; i++ {
		res += distance[i]
	}

	if total-res < res {
		res = total - res
	}
	return res
}

// @lc code=end

