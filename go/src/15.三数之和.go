/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	var length = len(nums)
	
	for i := 0; i < length; i++ {
		var j = i + 1
		var k = length - 1
		for j < k {
			var sum = nums[j] + nums[k]

			if sum == -nums[i] {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				k--
				j++
				for k > j && nums[k+1] == nums[k] {
					k--
				}
				for j < k && nums[j-1] == nums[j] {
					j++
				}

			} else if sum > -nums[i] {
				k--
				for k > j && nums[k+1] == nums[k] {
					k--
				}
			} else {
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			}
		}

		for i+1 < length && nums[i+1] == nums[i] {
			i++
		}
	}

	return res
}

// @lc code=end
