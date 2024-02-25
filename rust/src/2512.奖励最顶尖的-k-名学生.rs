/*
 * @lc app=leetcode.cn id=2512 lang=rust
 *
 * [2512] 奖励最顶尖的 K 名学生
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn top_students(
        positive_feedback: Vec<String>,
        negative_feedback: Vec<String>,
        report: Vec<String>,
        student_id: Vec<i32>,
        k: i32,
    ) -> Vec<i32> {
        let mut score = HashMap::new();
        for feedback in positive_feedback {
            score.insert(feedback, 3);
        }
        for feedback in negative_feedback {
            score.insert(feedback, -1);
        }

        let mut students = vec![];
        for (index, &id) in student_id.iter().enumerate() {
            let mut rank = 0;
            for word in report[index].split_whitespace() {
                if let Some(s) = score.get(word) {
                    rank += *s;
                }
            }

            students.push((id, rank));
        }

        students.sort_by(|a, b| match b.1.cmp(&a.1) {
            std::cmp::Ordering::Greater => std::cmp::Ordering::Greater,
            std::cmp::Ordering::Less => std::cmp::Ordering::Less,
            std::cmp::Ordering::Equal => a.0.cmp(&b.0),
        });

        let mut res = vec![];
        for i in 0..k as usize {
            res.push(students[i].0);
        }
        res
    }
}
// @lc code=end
