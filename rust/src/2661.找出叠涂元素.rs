/*
 * @lc app=leetcode.cn id=2661 lang=rust
 *
 * [2661] 找出叠涂元素
 */

// @lc code=start
impl Solution {
    pub fn first_complete_index(arr: Vec<i32>, mat: Vec<Vec<i32>>) -> i32 {
        let mut map = std::collections::HashMap::new();
        let (m, n) = (mat.len(), mat[0].len());
        for i in 0..m {
            for j in 0..n {
                map.insert(mat[i][j], (i, j));
            }
        }

        let (mut row, mut col) = (vec![0; m], vec![0; n]);
        for i in 0..arr.len() {
            let num = arr[i];
            let (r, c) = map.get(&num).unwrap();
            row[*r] += 1;
            col[*c] += 1;
            if row[*r] == n || col[*c] == m {
                return i as i32;
            }
        }

        -1
    }
}
// @lc code=end
