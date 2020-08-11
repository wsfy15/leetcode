/*
 * @lc app=leetcode.cn id=1290 lang=golang
 *
 * [1290] 二进制链表转整数
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res <<= 1
		res |= head.Val
		head = head.Next
	}
	return res
}

// @lc code=end

