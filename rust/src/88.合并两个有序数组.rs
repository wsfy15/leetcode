/*
 * @lc app=leetcode.cn id=88 lang=rust
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut cur = (m + n - 1) as usize;
        let mut m = m - 1;
        let mut n = n - 1;
        while (m >= 0 && n >= 0) {
            if (nums1[m as usize] > nums2[n as usize]) {
                nums1[cur] = nums1[m as usize];
                m -= 1;
            } else {
                nums1[cur] = nums2[n as usize];
                n -= 1;
            }
            cur -= 1;
        }

        while(n >= 0) {
            nums1[cur] = nums2[n as usize];
            cur -= 1;
            n -= 1;
        }
    }
}
// @lc code=end

