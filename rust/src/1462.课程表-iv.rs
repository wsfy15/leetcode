/*
 * @lc app=leetcode.cn id=1462 lang=rust
 *
 * [1462] 课程表 IV
 */

// @lc code=start
impl Solution {
    pub fn check_if_prerequisite(
        num_courses: i32,
        prerequisites: Vec<Vec<i32>>,
        queries: Vec<Vec<i32>>,
    ) -> Vec<bool> {
        let num_courses = num_courses as usize;
        let mut graph = vec![vec![false; num_courses]; num_courses];
        prerequisites
            .iter()
            .for_each(|edge| graph[edge[0] as usize][edge[1] as usize] = true);

        for k in 0..num_courses {
            for i in 0..num_courses {
                for j in 0..num_courses {
                    if graph[i][k] && graph[k][j] {
                        graph[i][j] = true;
                    }
                }
            }
        }

        let mut res = vec![];
        queries.iter().for_each(|query| {
            res.push(graph[query[0] as usize][query[1] as usize]);
        });

        res
    }
}
// @lc code=end
