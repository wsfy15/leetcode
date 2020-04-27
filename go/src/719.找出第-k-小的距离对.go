/*
 * @lc app=leetcode.cn id=719 lang=golang
 *
 * [719] 找出第 k 小的距离对
 */

// @lc code=start
// 1 <= k <= len(nums) * (len(nums) - 1) / 2 因此返回的是所有数对里面差值第k小的，而不是差值里面第k小的
// 差值为[0,0,0,1,2,3]，第1、2、3小的都是0
func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	low, high := 0, nums[n-1] - nums[0] // 对差值进行二分

	prefix := make([]int, nums[n-1] + 1) // prefix[i]：小于等于i的元素有多少个
	index := 0
	for i := 0; i <= nums[n-1]; i++ {
		for index < n && nums[index] == i {
			index++
		}
		prefix[i] = index // [0, index] 这些元素的值都小于等于index
	}

	multiplicity := make([]int, n) // multiplicity[i]：在nums[i]前面有多少个重复的
	for i := 1; i < n; i++ {
		if nums[i] == nums[i - 1] {
			multiplicity[i] = 1 + multiplicity[i - 1]
		}
	}

	for low < high {
		mid := low + (high - low) >> 1

		// 计算差值小于等于mid的数对 数量
		count := 0
		for i := 0; i < n; i++ {
			up := nums[i] + mid
			if up > nums[n-1] {
				up = nums[n-1]
			}

			// 每个数跟它后面的不大于up的数形成数对
			// 减去的multiplicity[i]可以这么理解：
			// 例如[..., 5, 5, 5, ...]，假设第一个5的下标为index，则multiplicity[index, index+1, index+2]=[0, 1, 2]
			// 当访问到index下标时，它可以跟后面的第二个和第三个5组成数对，但是prefix[5] - multiplicity[index] = prefix[5]
			// 此时，新增的count是大于5而小于等于up的元素的个数，并没有考虑到后面的两个5
			// 是在后面访问index+2下标时补偿回来的
			count += prefix[up] - (prefix[nums[i]] - multiplicity[i]) 
		}

		if count >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
// @lc code=end

