/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func oddEvenList(head *ListNode) *ListNode {
  if head == nil {
    return nil
  }
  odd, even := head, head.Next
  for even != nil && even.Next != nil {
    tmp := odd.Next
    odd.Next = even.Next
    even.Next = even.Next.Next
    odd.Next.Next = tmp
    odd, even = odd.Next, even.Next
  }
  
  return head
}
// @lc code=end

