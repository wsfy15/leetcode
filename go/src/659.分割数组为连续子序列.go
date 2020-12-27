/*
 * @lc app=leetcode.cn id=659 lang=golang
 *
 * [659] 分割数组为连续子序列
 */

// @lc code=start
func isPossible(nums []int) bool {
	n := len(nums)

	counter := make([]int, nums[n-1]-nums[0]+1)
	for i := n - 1; i >= 0; i-- {
		nums[i] -= nums[0]
		counter[nums[i]]++
	}

	tail := make([]int, len(counter))

	for _, num := range nums {
		if counter[num] == 0 {
			continue
		}

		if num > 2 && tail[num-1] > 0 {
			counter[num]--
			tail[num-1]--
			tail[num]++
		} else if num <= nums[n-1]-2 && counter[num+1] > 0 && counter[num+2] > 0 {
			counter[num]--
			counter[num+1]--
			counter[num+2]--
			tail[num+2]++
		} else {
			return false
		}
	}

	return true
}

// @lc code=end

