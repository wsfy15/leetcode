use std::cell::RefCell;
use std::collections::{BTreeMap, BinaryHeap, HashMap, HashSet, VecDeque};
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

fn main() {
    let edges = vec![[1, 1]];
    let edges: Vec<_> = edges.iter().map(|&arr| arr.to_vec()).collect();
    let mut res = edges;
    // let res = distinct_difference_array(vec![1, 2, 3, 4, 5]);

    // let edges: Vec<_> = edges.iter().map(|&arr| {arr.to_vec()}).collect();
    // let edges = vec![vec![1,2],vec![2,4],vec![1,3],vec![2,3],vec![2,1]];
    // let res = count_pairs(6, edges, vec![9]);
    // println!("{:?}", res);
    return;
}
