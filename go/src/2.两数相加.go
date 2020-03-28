/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  if l1 == nil {
    return l2
  }
  if l2 == nil {
    return l1
  }
  
  over := 0
  cur1, cur2 := l1, l2
  var prev *ListNode
  for cur1 != nil && cur2 != nil {
    cur1.Val = cur1.Val + cur2.Val + over
    if cur1.Val > 9 {
      cur1.Val, over = cur1.Val % 10, 1
    } else {
      over = 0
    }
    cur1, cur2, prev = cur1.Next, cur2.Next, cur1
  }
  
  left := cur1
  if cur1 == nil {
    left = cur2
  }
  
  if left != nil {
    prev.Next = left
    for over > 0 && left != nil {
      left.Val += over
      if left.Val > 9 {
        left.Val, over = 0, 1
      } else {
        over = 0
      }
      left, prev = left.Next, left
    }
  }
  
  if over > 0 && left == nil {
    prev.Next = &ListNode{
      Val: 1,
      Next: nil,
    }
  }

  return l1
}
// @lc code=end

