/*
 * @lc app=leetcode.cn id=1671 lang=rust
 *
 * [1671] 得到山形数组的最少删除次数
 */

// @lc code=start
impl Solution {
    pub fn minimum_mountain_removals(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let (mut left_del, mut right_del) = (vec![0; n], vec![0; n]);
        for i in 0..n {
            left_del[i] = i;
            right_del[i] = n - i - 1;
        }

        for i in 0..n {
            for j in 0..i {
                if nums[j] < nums[i] {
                    left_del[i] = left_del[i].min(left_del[j] + i - j - 1);
                }
            }
        }
        for i in (0..n).rev() {
            for j in (i + 1)..n {
                if nums[i] > nums[j] {
                    right_del[i] = right_del[i].min(right_del[j] + j - i - 1);
                }
            }
        }

        let mut res = i32::MAX;
        for i in 1..n {
            // 两端不能单调递减
            if left_del[i] == i || right_del[i] == n - i - 1 {
                continue;
            }
            res = res.min((left_del[i] + right_del[i]) as i32);
        }

        res
    }
}
// @lc code=end
