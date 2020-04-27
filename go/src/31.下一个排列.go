/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int)  {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i - 1] {
			minIndex := i
			for j := i + 1; j < n; j++ {
				if nums[j] > nums[i - 1] && nums[j] <= nums[minIndex] {
					minIndex = j
				}
			}

			nums[i - 1], nums[minIndex] = nums[minIndex], nums[i - 1]
			// 并不需要真正排序 nums[i: n]， 因为nums[i: n]已经是降序排序，直接反置一下就可以了
			for j := n - 1; j > i; {
				nums[i], nums[j] = nums[j], nums[i]
				j--
				i++
			}
			return
		}
	}

	for i, j := n - 1, 0; i > j; {
		nums[i], nums[j] = nums[j], nums[i]
		i--
		j++
	}
}

func nextPermutation2(nums []int)  {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i - 1] {
			minIndex := i
			for j := i + 1; j < n; j++ {
				if nums[j] > nums[i - 1] && nums[j] < nums[minIndex] {
					minIndex = j
				}
			}

			nums[i - 1], nums[minIndex] = nums[minIndex], nums[i - 1]
			// sort nums[i: n]
			for m := i; m < n; m++ {
				minIndex := m
				for k := m + 1; k < n; k++ {
					if nums[k] < nums[minIndex] {
						minIndex = k
					}
				}
				nums[m], nums[minIndex] = nums[minIndex], nums[m]
			}
			return
		}
	}

	for i, j := n - 1, 0; i > j; {
		nums[i], nums[j] = nums[j], nums[i]
		i--
		j++
	}
}
// @lc code=end

