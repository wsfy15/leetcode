/*
 * @lc app=leetcode.cn id=466 lang=rust
 *
 * [466] 统计重复个数
 */

// @lc code=start
use std::collections::HashMap;
impl Solution {
    pub fn get_max_repetitions(s1: String, n1: i32, s2: String, n2: i32) -> i32 {
        let s1 = s1.as_bytes();
        let s2 = s2.as_bytes();
        let (mut index, mut s1_cnt, mut s2_cnt) = (0, 0, 0);
        let (len1, len2) = (s1.len(), s2.len());

        // key为s2的index，value为当前的s1cnt, s2cnt
        let mut appear: HashMap<usize, Vec<i32>> = HashMap::new();

        while s1_cnt < n1 {
            s1_cnt += 1;
            for i in 0..len1 {
                if s1[i] == s2[index] {
                    index += 1;
                    if index == len2 {
                        index = 0;
                        s2_cnt += 1;
                    }
                }
            }

            // 标记s1扫描完后，之前是否也出现s2字符串判断到这个位置
            match appear.get(&index) {
                Some(v) => {
                    // 遍历了(s1cnt - v[0])个s1，发现有(s2cnt - v[1])个s1
                    let (s1_num, s2_num) = (s1_cnt - v[0], s2_cnt - v[1]);
                    let left = n1 - s1_cnt; // 还剩下多少个s1
                    let loops = left / s1_num; // 还能重复多少次
                    s1_cnt += loops * s1_num;
                    s2_cnt += loops * s2_num;
                }
                None => {
                    appear.insert(index, vec![s1_cnt, s2_cnt]);
                }
            }
        }

        while s1_cnt < n1 {
            s1_cnt += 1;
            for i in 0..len1 {
                if s1[i] == s2[index] {
                    index += 1;
                    if index == len2 {
                        index = 0;
                        s2_cnt += 1;
                    }
                }
            }
        }

        s2_cnt / n2
    }
}
// @lc code=end
