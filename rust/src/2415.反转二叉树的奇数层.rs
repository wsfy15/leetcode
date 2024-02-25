/*
 * @lc app=leetcode.cn id=2415 lang=rust
 *
 * [2415] 反转二叉树的奇数层
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
    pub fn reverse_odd_levels(
        root: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        let mut r = root.clone();
        let mut queue = vec![r];
        let mut level = 0;
        while !queue.is_empty() {
            let mut new_queue = vec![];
            for node in &queue {
                if let Some(n) = &node.clone().unwrap().borrow().left {
                    new_queue.push(Some(n.clone()));
                }
                if let Some(n) = &node.clone().unwrap().borrow().right {
                    new_queue.push(Some(n.clone()));
                }
            }

            if level % 2 == 1 {
                for i in 0..(queue.len() / 2) {
                    let j = queue.len() - i - 1;
                    let val = queue[j].clone().unwrap().borrow().val;
                    queue[j].clone().unwrap().borrow_mut().val =
                        queue[i].clone().unwrap().borrow().val;
                    queue[i].clone().unwrap().borrow_mut().val = val;
                }
            }
            level += 1;
            queue = new_queue;
        }

        root
    }
}
// @lc code=end
