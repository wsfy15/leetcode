/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
func splitArray(nums []int, m int) int {
	total, max := 0, 0
	for _, v := range nums {
		total += v
		if v > max {
			max = v
		}
	}

	// 最大值的最小值 介于 最大元素 和 nums和 之间，通过二分查找
	low, high := max, total 
	for low < high {
		// 以mid为最大值，所需数组数
		mid := low + (high - low) >> 1
		tmp, count := 0, 1
		for _, v := range nums {
			if tmp + v > mid {
				count++
				tmp = v

				if count > m {
					break
				}
			} else {
				tmp += v
			}
		}

		if count > m {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
// @lc code=end

