/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	for n > 1 {
		for i := 0; i < n / 2; i++ {
			lists[i] = merge(lists[i], lists[n - 1 - i])
		}

		if n & 1 == 1 {
			n++
		}
		n /= 2
	}

	return lists[0]
}

func merge(listA, listB *ListNode) *ListNode {
	var dummy ListNode
	dummy.Next = listA
	prev, cur := &dummy, dummy.Next
	curB := listB
	for curB != nil && cur != nil {
		if curB.Val < cur.Val {
			tmp := curB.Next
			prev.Next, curB.Next = curB, cur
			curB = tmp
		} else {
			cur = cur.Next
		}

		prev = prev.Next
	}

	if curB != nil {
		prev.Next = curB
	}

	return dummy.Next
}
// @lc code=end

