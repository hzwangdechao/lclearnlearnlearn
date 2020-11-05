/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (52.04%)
 * Likes:    483
 * Dislikes: 0
 * Total Accepted:    58.8K
 * Total Submissions: 110.5K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
 *
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
 *
 * 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
 *
 *
 *
 * 示例 1:
 *
 * 输入: 2, [[1,0]]
 * 输出: true
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
 *
 * 示例 2:
 *
 * 输入: 2, [[1,0],[0,1]]
 * 输出: false
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 * 1 <= numCourses <= 10^5
 *
 *
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	// return methodWithDfs(numCourses, prerequisites)
	return methodWithBfs(numCourses, prerequisites)
}

func methodWithDfs(numCourses int, prerequisites [][]int) bool  {
	var (
		edges = make([][]int, numCourses)
		visited = make([]int, numCourses)  // 0未搜索  1搜索中  2已经搜索完成
		result []int
		valid = true
		dfs func(u int)
	)
	dfs = func (u int)  {
		visited[u] = 1  // 将状态改变成正在搜索中

		for _, v := range edges[u]{
			if visited[v] == 0{
				// 如果是没有搜索的状态的话， 就搜索v
				dfs(v)
				if !valid{ // 如果成环的话则返回
					return
				}
			}else if visited[v] == 1{  // 如果是已经搜索的状态， 则返回
				valid = false
				return
			}
		}

		// 回溯完成， 将状态变成已经搜索的状态
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites{
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	fmt.Println(edges)

	for i:=0; i<numCourses && valid; i++{
		if visited[i] == 0{
			dfs(i)
		}
	}
	return valid
}

func methodWithBfs(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		inedg = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites{
		edges[info[1]] = append(edges[info[1]], info[0])
		inedg[info[0]] += 1
	}

	Q := make([]int, 0)
	for i:=0; i<len(inedg); i++{
		if inedg[i] == 0{
			Q = append(Q, i)
		}
	}

	for len(Q) > 0{
		u := Q[0]
		Q = Q[1:]

		result = append(result, u)

		for _, v := range edges[u]{
			inedg[v] --
			if inedg[v] == 0{
				Q = append(Q, v)
			}
		}
	}

	return len(result) == numCourses
}


// @lc code=end
