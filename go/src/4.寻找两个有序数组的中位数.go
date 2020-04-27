/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k := (m + n + 1) / 2
	if (m + n) & 1 == 1 {
		return float64(binarySearch(k, nums1, 0, m - 1, nums2, 0, n - 1))
	}
	return (float64(binarySearch(k, nums1, 0, m - 1, nums2, 0, n - 1) + binarySearch(k + 1, nums1, 0, m - 1, nums2, 0, n - 1))) / 2.0
}

func binarySearch(k int, nums1 []int, start1, end1 int, nums2 []int, start2, end2 int) int {
	len1, len2 := end1 - start1 + 1, end2 - start2 + 1 // 剩余元素个数
	if len1 < len2 { // 保证nums1剩余元素更多
		return binarySearch(k, nums2, start2, end2, nums1, start1, end1)
	}
	if len2 == 0 {
		return nums1[start1 + k - 1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	aPos, bPos := start1 + min(k / 2, len1) - 1, start2 + min(k / 2, len2) - 1
	if nums1[aPos] < nums2[bPos] {
		return binarySearch(k - (aPos - start1 + 1), nums1, aPos + 1, end1, nums2, start2, end2)
	} else {
		return binarySearch(k - (bPos - start2 + 1), nums1, start1, end1, nums2, bPos + 1, end2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// 中位数是 第 ((m + n + 1) / 2 + (m + n + 2) / 2) 个数，下标从1开始
	// 如果m+n为奇数，那么上面两部分都是同一个值
	// 如果m+n为偶数，那么上面两部分就是 (m + n) / 2、1 + (m + n) / 2 这两个位置

	// O(m+n)的解法，双指针从两数组起始位置开始遍历，直到遍历了  ((m + n + 1) / 2 + (m + n + 2) / 2) 个数
	aPos, bPos := 0, 0
	last, cur := 0, 0
	// 直接遍历掉前1 + (m + n) / 2个数，则对于奇数长度来说，最后一个数为结果；对于偶数长度来说，最后两个数的平均值为结果
	for i := 0; i < 1 + (m + n) / 2; i++ {
		last = cur
		if aPos < m && (bPos >= n || nums1[aPos] < nums2[bPos]) {
			cur = nums1[aPos]
			aPos++
		} else {
			cur = nums2[bPos]
			bPos++
		}
	}

	if (m + n) & 1 == 1 {
		return float64(cur)
	}

	return float64(cur + last) / 2.0
}
// @lc code=end

