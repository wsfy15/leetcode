/*
 * @lc app=leetcode.cn id=765 lang=rust
 *
 * [765] 情侣牵手
 */

// @lc code=start
impl Solution {
    pub fn min_swaps_couples(row: Vec<i32>) -> i32 {
        let n = row.len();

        // 目标有m个联通集
        let mut m = n / 2;
        let mut uf = vec![0; m];
        for i in 0..m {
            uf[i] = i;
        }

        for i in (0..n).step_by(2) {
            // 如果需要做交换才能得到联通集，m减一
            if Self::merge((row[i] / 2) as usize, (row[i + 1] / 2) as usize, &mut uf) {
                m -= 1;
            }
        }

        (n / 2 - m) as i32
    }

    fn merge(i: usize, j: usize, uf: &mut Vec<usize>) -> bool {
        let (a, b) = (Self::find(i, uf), Self::find(j, uf));
        uf[a] = b;
        a != b
    }

    fn find(i: usize, uf: &mut Vec<usize>) -> usize {
        let mut i = i;
        while uf[i] != i {
            uf[i] = uf[uf[i]];
            i = uf[i];
        }
        i
    }
}
// @lc code=end
