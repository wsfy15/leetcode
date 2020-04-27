/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
// 双指针：排序后用两个指针去遍历

func intersect(nums1 []int, nums2 []int) []int {
  if len(nums1) < len(nums2) {
    nums1, nums2 = nums2, nums1
  }
  
  m1 := make(map[int]int)
  res := []int{}
  
  for _, v := range nums1 {
    m1[v]++
  }
  
  for _, v := range nums2 {
    if count, ok := m1[v]; ok && count > 0 {
      m1[v]--
      res = append(res, v)
    }
  }
  
  return res
}
// @lc code=end

