/*
 * @lc app=leetcode.cn id=56 lang=rust
 *
 * [56] 合并区间
 */

// @lc code=start
impl Solution {
    pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut intervals = intervals;
        intervals.sort();
        let mut res = vec![];
        let (mut start, mut end) = (intervals[0][0], intervals[0][1]);
        intervals.iter().skip(1).for_each(|x| {
            if x[0] > end {
                res.push(vec![start, end]);
                start = x[0];
            }
            end = end.max(x[1]);
        });

        res.push(vec![start, end]);

        res
    }

    pub fn merge2(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut intervals = intervals.clone();
        intervals.sort_by(|a, b| {
            if a[0] == b[0] {
                b[1].cmp(&a[1])
            } else {
                a[0].cmp(&b[0])
            }
        });

        let mut res = vec![];
        let mut startIndex = 0;
        let mut right = intervals[0][1];
        let n = intervals.len();
        for i in 1..n {
            if intervals[i][0] <= right {
                right = intervals[i][1].max(right);
            } else {
                res.push(vec![intervals[startIndex][0], right]);
                right = intervals[i][1];
                startIndex = i;
            }
        }
        res.push(vec![intervals[startIndex][0], right]);

        res
    }
}
// @lc code=end
