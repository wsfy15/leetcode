/*
 * @lc app=leetcode.cn id=103 lang=rust
 *
 * [103] 二叉树的锯齿形层序遍历
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
    pub fn zigzag_level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        if root.is_none() {
            return vec![];
        }

        let mut res = vec![];
        let mut queue = vec![root.clone().unwrap()];
        let mut l2r = true;
        while queue.len() > 0 {
            let mut new_queue = vec![];
            let mut cur = vec![];
            for i in 0..queue.len() {
                cur.push(queue[i].borrow().val);

                if let Some(left) = queue[i].borrow().left.clone() {
                    new_queue.push(left);
                }
                if let Some(right) = queue[i].borrow().right.clone() {
                    new_queue.push(right);
                }
            }

            if !l2r {
                cur.reverse();
            }

            res.push(cur);
            queue = new_queue;
            l2r = !l2r;
        }

        res
    }
}
// @lc code=end
