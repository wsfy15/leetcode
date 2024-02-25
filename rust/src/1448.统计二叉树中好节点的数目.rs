/*
 * @lc app=leetcode.cn id=1448 lang=rust
 *
 * [1448] 统计二叉树中好节点的数目
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
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn good_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        fn dfs(cur: Option<Rc<RefCell<TreeNode>>>, max: i32) -> i32 {
            if let Some(cur) = cur {
                let mut res = 0;
                let mut max = max;
                let curVal = cur.borrow().val;
                if curVal >= max {
                    res += 1;
                    max = curVal;
                }
                res += dfs(cur.borrow_mut().left.take(), max);
                res += dfs(cur.borrow_mut().right.take(), max);
                return res;
            }

            0
        }

        dfs(root, -100000)
    }
}
// @lc code=end

