/*
 * @lc app=leetcode.cn id=2807 lang=rust
 *
 * [2807] 在链表中插入最大公约数
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
    fn gcd(mut a: i32, mut b: i32) -> i32 {
        while a != 0 {
            (a, b) = (b % a, a);
        }
        b
    }

    pub fn insert_greatest_common_divisors(
        mut head: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut cur = &mut head;
        while cur.as_ref().unwrap().next.is_some() {
            let x = cur.as_mut().unwrap();
            let next = x.next.take();
            x.next = Some(Box::new(ListNode {
                val: Self::gcd(x.val, next.as_ref().unwrap().val),
                next,
            }));
            cur = &mut cur.as_mut().unwrap().next.as_mut().unwrap().next;
        }

        head
    }
}
// @lc code=end
