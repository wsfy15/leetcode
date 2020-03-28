/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
  // 指针A走完链表A，再从链表B走
  // 指针B走完链表B，再从链表A走
  // 由于两个指针前进速度一样，当到达交点时，走过的路程都一样
  if headA == nil || headB == nil {
    return nil
  }
  ptrA, ptrB := headA, headB
  for ptrA != nil || ptrB != nil { // 都走了A和B链表一遍还没交点的话，同时为nil
    if ptrA == nil {
      ptrA = headB
    }
    
    if ptrB == nil {
      ptrB = headA
    }
    
    if ptrA == ptrB {
      return ptrA
    }
    
    ptrA = ptrA.Next
    ptrB = ptrB.Next
  }
  return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
  // 求链表长度，然后让长的先走。一样长再一起走
  lenA, lenB := 0, 0
  cur := headA
  for cur != nil {
    lenA++
    cur = cur.Next
  }
  
  cur = headB
  for cur != nil {
    lenB++
    cur = cur.Next
  }
  
  for lenA > lenB {
    headA = headA.Next
    lenA--
  }
  
  for lenA < lenB {
    headB = headB.Next
    lenA++
  }
  
  for headA != headB {
    headA = headA.Next
    headB = headB.Next
  }
  return headA
}
// @lc code=end

