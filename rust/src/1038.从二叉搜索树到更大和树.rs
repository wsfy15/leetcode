/*
 * @lc app=leetcode.cn id=1038 lang=rust
 *
 * [1038] 从二叉搜索树到更大和树
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
    pub fn bst_to_gst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        Self::post_order(&root, &mut 0);
        root
    }

    fn post_order(root: &Option<Rc<RefCell<TreeNode>>>, sum: &mut i32) {
        if let Some(t) = root {
            Self::post_order(&t.borrow().right, sum);
            *sum += t.borrow_mut().val;
            t.borrow_mut().val = *sum;
            Self::post_order(&t.borrow().left, sum);
        }
    }
}
// @lc code=end
