/*
 * @lc app=leetcode.cn id=57 lang=rust
 *
 * [57] 插入区间
 */

// @lc code=start
impl Solution {
    pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        if new_interval.is_empty() {
            return intervals;
        }
        if intervals.is_empty() {
            return vec![new_interval];
        }

        let mut res = vec![];
        let n = intervals.len();
        let (mut start, mut end) = (new_interval[0], new_interval[1]);
        if end < intervals[0][0] {
            res.push(vec![start, end]);
            intervals.iter().for_each(|x| res.push(x.clone()));
            return res;
        } else if start > intervals[n - 1][1] {
            intervals.iter().for_each(|x| res.push(x.clone()));
            res.push(vec![start, end]);
            return res;
        }

        let mut index = 0;
        while index < n {
            if start <= intervals[index][1] {
                for i in 0..index {
                    res.push(intervals[i].clone());
                }

                if end < intervals[index][0] {
                    res.push(vec![start, end]);
                    break;
                }

                if start > intervals[index][0] {
                    start = intervals[index][0];
                }
                if end < intervals[index][1] {
                    end = intervals[index][1];
                    index += 1;
                } else {
                    index += 1;
                    while index < n && end >= intervals[index][0] {
                        index += 1;
                    }
                    if end < intervals[index - 1][1] {
                        end = intervals[index - 1][1];
                    }
                }
                res.push(vec![start, end]);
                break;
            }

            index += 1;
        }

        while index < n {
            res.push(intervals[index].clone());
            index += 1;
        }

        res
    }
}
// @lc code=end
