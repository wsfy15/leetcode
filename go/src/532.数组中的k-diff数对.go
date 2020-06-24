/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的K-diff数对
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	visited := map[int]bool{}
	result := map[int]bool{} // 记录pair中的较小值
	for _, num := range nums {
		if _, ok := visited[num - k]; ok {
			result[num - k] = true
		}
		if _, ok := visited[num + k]; ok {
			result[num] = true
		}
		visited[num] = true
	}
	return len(result)
}

func findPairs2(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	n := len(nums)
	if n <= 1 {
		return 0
	}

	res := 0
	sort.Ints(nums)
	i, j := 0, 1
	for j < n {
		diff := nums[j] - nums[i] 
		if diff <= k {
			if diff == k {
				res++
			}
			for j + 1 < n && nums[j + 1] == nums[j] {
				j++
			}
			j++
		} else {
			i++
			if i == j {
				j++
			}
		}
	}

	return res
}
// @lc code=end

