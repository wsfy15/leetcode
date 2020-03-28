/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
  if head == nil {
    return head
  }
  
  dummy := &ListNode{
    Next: head,
  }
  slow, fast := dummy, dummy.Next
  
  // 如果k是链长的很多倍，将很耗时，改用先求链长   
  // for k > 0 {
  //   fast = fast.Next
  //   if fast == nil {
  //     fast = dummy.Next
  //   }
  //   k--
  // }
  
  count := 0
  for fast != nil {
    fast = fast.Next
    count++
  }
  
  k %= count
  if k == 0 {
    return dummy.Next
  }
  
  fast = dummy
  for k > 0 {
    fast = fast.Next
    k--
  }
  
  for fast.Next != nil { 
    slow = slow.Next
    fast = fast.Next
  }
  fast.Next = dummy.Next
  dummy.Next = slow.Next
  slow.Next = nil
  
  return dummy.Next
}
// @lc code=end

