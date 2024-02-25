/*
 * @lc app=leetcode.cn id=145 lang=rust
 *
 * [145] 二叉树的后序遍历
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
use std::collections::VecDeque;
use std::rc::Rc;
impl Solution {
    pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        if root.is_none() {
            return vec![];
        }

        let mut res = VecDeque::new();
        let mut stack = VecDeque::new();
        stack.push_back(root.clone().unwrap());
        while stack.len() > 0 {
            let cur = stack.pop_back().unwrap();
            res.push_front(cur.borrow().val);

            if cur.borrow().left.is_some() {
                stack.push_back(cur.borrow().left.clone().unwrap());
            }
            if cur.borrow().right.is_some() {
                stack.push_back(cur.borrow().right.clone().unwrap());
            }
        }

        res.into()
    }
}
// @lc code=end
