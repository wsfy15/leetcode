/*
 * @lc app=leetcode.cn id=2487 lang=rust
 *
 * [2487] 从链表中移除节点
 */

// @lc code=start
// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
//
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn remove_nodes(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head.unwrap();
        if head.next.is_none() {
            return Some(head);
        }

        let node = Self::remove_nodes(head.next).unwrap();
        if node.val > head.val {
            // 返回的链表头一定是最大的
            return Some(node);
        }

        head.next = Some(node);
        return Some(head);
    }
}
// @lc code=end
