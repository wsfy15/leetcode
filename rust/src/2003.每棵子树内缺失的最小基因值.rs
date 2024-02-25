/*
 * @lc app=leetcode.cn id=2003 lang=rust
 *
 * [2003] 每棵子树内缺失的最小基因值
 */

// @lc code=start

use std::collections::BTreeSet;

impl Solution {
    pub fn smallest_missing_value_subtree(parents: Vec<i32>, nums: Vec<i32>) -> Vec<i32> {
        // 记录每个节点的子节点
        let n = parents.len();
        let e = {
            let mut e = vec![vec![]; n];
            for (i, &p) in parents.iter().enumerate().skip(1) {
                e[p as usize].push(i);
            }
            e
        };

        // 找到基因值为1的节点，从该节点到根节点的基因值都大于1，其他节点基因值均为1
        let mut ret = vec![1; n];
        let Some(one) = nums.iter().position(|&x| x == 1)
            else { return ret };

        // 通过 B tree 实现有序，将包含基因值1的子树中每个节点的基因值从B tree中清除
        let mut mark = (1..=100_001).collect();
        let arg = (&e, &nums);
        fn remove_mark(
            arg @ &(e, nums): &(&Vec<Vec<usize>>, &Vec<i32>),
            mark: &mut BTreeSet<i32>,
            node: usize,
        ) {
            mark.remove(&nums[node]);
            for &i in &e[node] {
                remove_mark(arg, mark, i);
            }
        }

        remove_mark(&arg, &mut mark, one);
        let mut node = one;
        loop {
            ret[node] = *mark.first().unwrap();
            if parents[node] == -1 {
                break;
            }
            let par = parents[node] as usize;
            for &sbl in &e[par] {
                if sbl == node {
                    continue;
                }
                remove_mark(&arg, &mut mark, sbl);
            }
            mark.remove(&nums[par]);
            node = par;
        }
        ret
    }
}
// @lc code=end
