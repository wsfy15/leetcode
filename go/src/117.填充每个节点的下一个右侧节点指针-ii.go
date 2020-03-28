/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
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
  if root == nil {
    return nil
  }
  
  leftmost := root
  for leftmost != nil {
    head := leftmost
    leftmost = nil
    var prev *Node = nil
    for head != nil {
      process(head.Left, &prev, &leftmost)
      process(head.Right, &prev, &leftmost)
      head = head.Next
    }
  }
  
  return root
}

// 从父层的下一个节点开始找子层节点
func process(cur *Node, prev, leftmost **Node) {
  if cur == nil {
    return
  }
  
  if *prev != nil {
    (*prev).Next = cur
  } else {
    *leftmost = cur
  }
  *prev = cur
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
    }
    
    if cur.Right != nil {
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

