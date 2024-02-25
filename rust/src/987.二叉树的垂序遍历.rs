/*
 * @lc app=leetcode.cn id=987 lang=rust
 *
 * [987] 二叉树的垂序遍历
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
use std::collections::BTreeMap;
use std::rc::Rc;
impl Solution {
    pub fn vertical_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        if root.is_none() {
            return vec![];
        }

        let mut stack = vec![(root.unwrap(), 0, 0)];
        let mut mp = BTreeMap::new();
        let mut tmp = vec![];
        while stack.len() > 0 {
            for i in 0..stack.len() {
                mp.entry(stack[i].1).or_insert(vec![]).push((
                    stack[i].1, // col
                    stack[i].2, // row
                    stack[i].0.borrow().val,
                ));

                if let Some(left) = stack[i].0.borrow_mut().left.take() {
                    tmp.push((left, stack[i].1 - 1, stack[i].2 + 1));
                }
                if let Some(right) = stack[i].0.borrow_mut().right.take() {
                    tmp.push((right, stack[i].1 + 1, stack[i].2 + 1));
                }
            }

            stack = tmp;
            tmp = vec![];
        }

        mp.into_values()
            .map(|mut v| {
                v.sort_unstable();
                v.into_iter().map(|item| item.2).collect::<Vec<i32>>()
            })
            .collect::<Vec<_>>()
    }
}
// @lc code=end
