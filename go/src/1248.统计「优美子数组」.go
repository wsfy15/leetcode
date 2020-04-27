/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] 统计「优美子数组」
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	odds := []int{-1} // 记录奇数下标
	for i, v := range nums {
		if v & 1 == 1 {
			odds = append(odds, i)
		}
	}
	odds = append(odds, len(nums))

	res := 0
	for i := 1; i + k < len(odds); i++ {
		res += (odds[i] - odds[i - 1]) * (odds[i + k] - odds[i + k - 1])
	}
	return res
}

func numberOfSubarrays3(nums []int, k int) int {
	odds := make([]int, len(nums) + 1) // odds[i]：nums[0, j]，当前有i个奇数的j的个数
	odds[0] = 1
	odd := 0 // 当前有多少个奇数
	res := 0
	for _, v := range nums {
		if v & 1 == 1 {
			odd++
		}
		odds[odd]++
		if odd >= k {
			res += odds[odd - k] // 前面（包含自己）有odd-k个奇数的元素,到当前位置共有k个奇数
		}
	}
	return res
}

func numberOfSubarrays2(nums []int, k int) int {
	i, j := 0, 0
	cur, res := 0, 0
	n := len(nums)

	for j < n {
		if nums[j] & 1 == 1 {
			cur++
			if cur == k {
				tmp := j + 1
				for tmp < n && nums[tmp] & 1 == 0 {
					tmp++
				}
				postfix := tmp - j

				for cur == k {
					res += postfix
					if nums[i] & 1 == 1 {
						cur--
					} 
					i++
				}
			}
		}
		j++
	}

	return res
}
// @lc code=end

