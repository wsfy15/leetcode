/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
  left, right := 0, len(nums) - 1
  for left <= right {
    mid := left + (right - left) >> 1
    if nums[mid] == target {
      return mid
    }
    
    if nums[left] <= nums[mid] { // 左边升序
      if nums[left] <= target && target <= nums[mid] {
        right = mid - 1
      } else {
        left = mid + 1
      }
    } else { // 右边升序
      if nums[mid] <= target && target <= nums[right] {
        left = mid + 1
      } else {
        right = mid - 1
      }
    }
  }
  
  return -1
}

func search2(nums []int, target int) int {
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

