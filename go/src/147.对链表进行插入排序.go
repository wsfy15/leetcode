/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{ Next: head }
	curPrev, cur := dummy, head
	for cur != nil {
		prev, tmp := dummy, dummy.Next
		for tmp != cur {
			if cur.Val < tmp.Val {
				curPrev.Next = cur.Next
				cur.Next = prev.Next
				prev.Next = cur
				break
			} else {
				tmp, prev = tmp.Next, prev.Next
			}
		}

		if curPrev.Next == cur {
			cur, curPrev = cur.Next, curPrev.Next
		} else {
			cur = curPrev.Next
		}
	}

	return dummy.Next
}
// @lc code=end

