/*
 * @lc app=leetcode.cn id=2583 lang=rust
 *
 * [2583] 二叉树中的第 K 大层和
 */

// @lc code=start
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn kth_largest_level_sum(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i64 {
        if root.is_none() {
            return -1;
        }

        let mut sum = vec![];
        let mut queue = vec![root.unwrap()];

        while queue.len() > 0 {
            let mut val = 0i64;
            let mut new_queue = vec![];
            for node in &queue {
                val += node.borrow().val as i64;
                if let Some(left) = node.borrow().left.clone() {
                    new_queue.push(left);
                }
                if let Some(right) = node.borrow().right.clone() {
                    new_queue.push(right);
                }
            }

            sum.push(val);
            queue = new_queue;
        }

        if sum.len() < k as usize {
            return -1;
        }

        sum.sort_by(|a, b| b.cmp(a));
        sum[k as usize - 1]
    }
}
// @lc code=end
