/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  另一种解法：找中点，翻转后半部分，依次对比
 func isPalindrome(head *ListNode) bool {
  if head == nil {
    return true
  }
  cur := head
  var vals []int
  for cur != nil {
    vals = append(vals, cur.Val)
    cur = cur.Next
  }
  
  i, j := 0, len(vals) - 1
  for i < j {
    if vals[i] == vals[j] {
      i++
      j--
    } else {
      return false
    }
  }
  return true
}
// @lc code=end

