/*
 * @lc app=leetcode.cn id=2641 lang=rust
 *
 * [2641] 二叉树的堂兄弟节点 II
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
    pub fn replace_value_in_tree(
        root: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return None;
        }

        root.clone().unwrap().borrow_mut().val = 0;
        let mut queue = vec![root.clone().unwrap()];

        while queue.len() > 0 {
            let current = queue;
            queue = vec![];

            let mut sum = 0;
            current.iter().for_each(|node| {
                if let Some(child) = node.borrow().left.clone() {
                    sum += child.borrow().val;
                    queue.push(child);
                }
                if let Some(child) = node.borrow().right.clone() {
                    sum += child.borrow().val;
                    queue.push(child);
                }
            });

            current.iter().for_each(|node| {
                let mut child_sum = 0;
                if let Some(child) = node.borrow().left.clone() {
                    child_sum += child.borrow().val;
                }
                if let Some(child) = node.borrow().right.clone() {
                    child_sum += child.borrow().val;
                }

                if let Some(child) = node.borrow().left.clone() {
                    child.borrow_mut().val = sum - child_sum;
                }
                if let Some(child) = node.borrow().right.clone() {
                    child.borrow_mut().val = sum - child_sum;
                }
            })
        }

        root
    }
}
// @lc code=end
