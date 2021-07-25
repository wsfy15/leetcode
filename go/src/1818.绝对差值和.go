/*
 * @lc app=leetcode.cn id=1818 lang=golang
 *
 * [1818] 绝对差值和
 */

// @lc code=start
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n, res := len(nums1), 0
	var mod int = 1e9 + 7
	set := map[int]struct{}{}
	for i := 0; i < n; i++ {
		res += abs(nums1[i] - nums2[i])
		set[nums1[i]] = struct{}{}
	}

	nums := []int{}
	for k, _ := range set {
		nums = append(nums, k)
	}
	sort.Ints(nums)

	diff := 0
	for i := 0; i < n; i++ {
		// 找到最接近nums[2]的元素
		index := lowerBound(nums, nums2[i])
		origin := abs(nums2[i] - nums1[i])
		diff = max(diff, origin-abs(nums2[i]-nums[index]))
		if index > 0 {
			diff = max(diff, origin-abs(nums2[i]-nums[index-1]))
		}
		if index < len(nums)-1 {
			diff = max(diff, origin-abs(nums2[i]-nums[index+1]))
		}
	}

	return (res - diff) % mod
}

// 找到nums中大于等于target的元素的index
func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}

// @lc code=end

