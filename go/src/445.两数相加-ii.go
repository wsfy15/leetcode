/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []int{}, []int{}
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	var dummy ListNode
	carry := 0
	for len(stack1) > 0 && len(stack2) > 0 {
		tmp := stack1[len(stack1) - 1] + stack2[len(stack2) - 1] + carry
		if tmp >= 10 {
			carry = 1
			tmp -= 10
		} else {
			carry = 0
		}

		node := &ListNode{
			Val: tmp,
			Next: dummy.Next,
		}
		dummy.Next = node

		stack1 = stack1[:len(stack1) - 1]
		stack2 = stack2[:len(stack2) - 1]
	}

	stack := stack1
	if len(stack2) > 0 {
		stack = stack2
	}

	for len(stack) > 0 {
		tmp := stack[len(stack) - 1] + carry
		if tmp >= 10 {
			carry = 1
			tmp -= 10
		} else {
			carry = 0
		}

		node := &ListNode{
			Val: tmp,
			Next: dummy.Next,
		}
		dummy.Next = node

		stack = stack[:len(stack) - 1]
	}

	if carry > 0 {
		dummy.Next = &ListNode{
			Val: carry,
			Next: dummy.Next,
		}
	}

	return dummy.Next
}
// @lc code=end

