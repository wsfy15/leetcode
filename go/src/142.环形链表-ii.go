/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
  count := make(map[*ListNode]int)
  index := 0
  for head != nil {
    if _, ok := count[head]; ok {
      return head
    }
    count[head] = index
    index++
    head = head.Next
  }
  return nil
}
// @lc code=end

