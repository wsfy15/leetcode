/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
		n := len(nums)
		if n == 0 {
			return -1
		}
		
		i := 0
		for ; i < n - 1; i++ {
			if nums[i] > nums[i + 1] {
				break;
			}
		}

		var low, high int
		if nums[0] <= target && nums[i] >= target {
			low = 0
			high = i
		} else {
			low = i + 1
			high = n - 1
		}

		for low <= high {
			mid := low + ((high - low) >> 1)
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		return -1
}
// @lc code=end

