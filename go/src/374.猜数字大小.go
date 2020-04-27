/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low, high := 1, n
	for low < high {
		mid := low + (high - low) >> 1

		tmp := guess(mid)
		if tmp == 0 {
			return mid
		} else if tmp == -1 {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
// @lc code=end

