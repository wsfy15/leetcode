/*
 * @lc app=leetcode.cn id=210 lang=rust
 *
 * [210] 课程表 II
 */

// @lc code=start
impl Solution {
    pub fn find_order(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
        let mut order = vec![];
        let mut in_degree = vec![0; num_courses as usize];

        // 记录每个课程的后续课程
        let mut edges = vec![vec![]; num_courses as usize];
        prerequisites.iter().for_each(|edge| {
            in_degree[edge[0] as usize] += 1;
            edges[edge[1] as usize].push(edge[0]);
        });

        let mut queue: Vec<_> = in_degree
            .iter()
            .enumerate()
            .filter(|(_, &degree)| degree == 0)
            .map(|(i, _)| i)
            .collect();
        let mut finish_courses = 0;
        while queue.len() > 0 {
            let index = queue.remove(0);
            finish_courses += 1;
            order.push(index as i32);

            edges[index].iter().for_each(|&node| {
                in_degree[node as usize] -= 1;
                if in_degree[node as usize] == 0 {
                    queue.push(node as usize);
                }
            });
        }

        if finish_courses == num_courses {
            order
        } else {
            vec![]
        }
    }
}
// @lc code=end
