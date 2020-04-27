/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	for i := 0; i < n - 3; i++ {
		for j := i + 1; j < n - 2; j++ {
			left, right := j + 1, n - 1
			for left < right {
				cur := nums[i] + nums[j] + nums[left] + nums[right] - target
				if cur == 0 {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left + 1] {
						left++
					}
					for left < right && nums[right] == nums[right - 1] {
						right--
					}
					left++
					right--
				} else if cur > 0 {
					right--
				} else {
					left++
				}
			
			}

			for j < n - 1 && nums[j] == nums[j + 1] {
				j++
			}
		}

		for i < n - 1 && nums[i] == nums[i + 1] {
			i++
		}
	}

	return res
}
// @lc code=end

