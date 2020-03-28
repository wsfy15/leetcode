/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
 func connect(root *Node) *Node {
  // 利用已建立的next指针
  if root == nil {
    return nil
  }
  
  leftmost := root
  for leftmost.Left != nil {
    head := leftmost
    for head != nil {
      head.Left.Next = head.Right
      if head.Next != nil {
        head.Right.Next = head.Next.Left
      } else {
        head.Right.Next = nil
      }
      
      head = head.Next
    }
    leftmost = leftmost.Left
  }
  return root
}

func connect2(root *Node) *Node {
  if root == nil {
    return nil
  }
  
  queue := []*Node{root}
  backup := []*Node{}
  for len(queue) > 0 {
    cur := queue[0]
    queue = queue[1:]
    
    if cur.Left != nil {
      backup = append(backup, cur.Left)
      backup = append(backup, cur.Right)
    }
    
    if len(queue) > 0 {
      cur.Next = queue[0]
    } else {
      cur.Next = nil
      queue, backup = backup, queue
    }
  }
  
  return root
}
// @lc code=end

