/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
  n2n := make(map[*Node]*Node)
  dummy := new(Node)
  cur, ptr := dummy, head
  for ptr != nil {
    cur.Next = &Node{
      Val: ptr.Val,
    }
    n2n[ptr] = cur.Next
    cur, ptr = cur.Next, ptr.Next
  }
  
  cur, ptr = dummy.Next, head
  for ptr != nil {
    if r, ok := n2n[ptr.Random]; ok {
      cur.Random = r
    }
    cur, ptr = cur.Next, ptr.Next
  }
  
  
  return dummy.Next
}
// @lc code=end

