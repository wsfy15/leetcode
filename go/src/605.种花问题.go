/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	zeros, count := 0, 0
	if flowerbed[0] == 0 {
		zeros++
	}
	for _, v := range flowerbed {
		if v == 0 {
			zeros++
		} else {
			if zeros > 0 {
				count += (zeros - 1) >> 1
			}
			zeros = 0
		}
	}
	if zeros > 0 {
		count += zeros >> 1
	}

	return count >= n
}
// @lc code=end

