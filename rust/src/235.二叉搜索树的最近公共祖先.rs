/*
 * @lc app=leetcode.cn id=235 lang=rust
 *
 * [235] 二叉搜索树的最近公共祖先
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
    pub fn lowest_common_ancestor(
        root: Option<Rc<RefCell<TreeNode>>>,
        p: Option<Rc<RefCell<TreeNode>>>,
        q: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        let pv = p.clone().unwrap().borrow().val;
        let qv = q.clone().unwrap().borrow().val;
        let val = root.clone().unwrap().borrow().val;
        if (pv <= val && val <= qv) || (qv <= val && val <= pv) {
            return root;
        }

        if pv < val {
            return Self::lowest_common_ancestor(root.unwrap().borrow().left.clone(), p, q);
        }
        return Self::lowest_common_ancestor(root.unwrap().borrow().right.clone(), p, q);
    }
}
// @lc code=end
