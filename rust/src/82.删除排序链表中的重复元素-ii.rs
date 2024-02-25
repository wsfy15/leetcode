/*
 * @lc app=leetcode.cn id=82 lang=rust
 *
 * [82] 删除排序链表中的重复元素 II
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
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut dummy = ListNode::new(0);
        let (mut prev, mut cur, mut cnt) = (&mut dummy, &head, 0);

        while let Some(cur_node) = cur {
            let next = &cur_node.next;

            if let Some(next_node) = next {
                if next_node.val == cur_node.val {
                    cnt += 1;
                } else {
                    if cnt <= 0 {
                        let node = ListNode::new(cur_node.val);

                        prev.next = Some(Box::new(node));
                        prev = prev.next.as_mut().unwrap();
                    }
                    cnt = 0;
                }
            } else {
                if cnt <= 0 {
                    let node = ListNode::new(cur_node.val);

                    prev.next = Some(Box::new(node));
                }
                break;
            }

            cur = next;
        }

        dummy.next
    }
}
// @lc code=end
