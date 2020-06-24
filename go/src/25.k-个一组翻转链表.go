/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	var dummy ListNode
	dummy.Next = head
	cur := dummy.Next
	prev := &dummy
	for cur != nil {
		tmp := cur
		count := 1

		for count < k && tmp.Next != nil {
			tmp = tmp.Next
			count++
		}

		if count != k {
			return dummy.Next
		}

		next := tmp.Next
		last := tmp
		curBackup := cur
		// 依次把每个节点放到last后面去
		for i := 1; i < k; i++ {
			prev.Next = cur.Next
			cur.Next = last.Next
			last.Next = cur
			cur = prev.Next
		}

		cur = next
		prev = curBackup
	}

	return dummy.Next
}
// @lc code=end

