/*
 * @lc app=leetcode.cn id=83 lang=rust
 *
 * [83] 删除排序链表中的重复元素
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
    pub fn delete_duplicates(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        match head {
            Some(cur) => match cur.next {
                Some(next_node) => {
                    if next_node.val == cur.val {
                        Self::delete_duplicates(Some(next_node))
                    } else {
                        let mut node = ListNode::new(cur.val);
                        node.next = Self::delete_duplicates(Some(next_node));
                        Some(Box::new(node))
                    }
                }
                None => Some(cur),
            },
            None => None,
        }
    }
}
// @lc code=end
