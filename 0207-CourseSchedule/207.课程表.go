/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	// solution1: DFS
	courseRequisites := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		courseRequisites[prerequisite[0]] = append(courseRequisites[prerequisite[0]], prerequisite[1])
	}
	// 0表示待学习，1表示学习中，2表示学习完成
	memory := make([]int, numCourses)

	var dfs func(courseIdx int) bool
	dfs = func(courseIdx int) bool {
		memory[courseIdx] = 1
		for _, nextCourse := range courseRequisites[courseIdx] {
			if memory[nextCourse] == 0 {
				if !dfs(nextCourse) {
					return false
				}
			} else if memory[nextCourse] == 1 {
				return false
			}
		}
		memory[courseIdx] = 2
		return true
	}
	for courseIdx := 0; courseIdx < numCourses; courseIdx += 1 {
		if !dfs(courseIdx) {
			return false
		}
	}
	return true
	// solution2: 广度优先搜索
	// todo
}

// @lc code=end

