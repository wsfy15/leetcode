/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
// 双指针：排序两个数组后，用两个指针去遍历
// 二分：排序其中一个数组，然后在该数组上用二分查找搜索另一个数组的每个值

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	for _, v := range nums1 {
		m[v] = true
	}

	res := []int{}
	for _, v := range nums2 {
		if flag, ok := m[v]; ok && flag {
			res = append(res, v)
			m[v] = false
		}
	}
	return res
}
// @lc code=end

