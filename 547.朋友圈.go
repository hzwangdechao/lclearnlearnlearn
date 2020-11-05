/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 *
 * https://leetcode-cn.com/problems/friend-circles/description/
 *
 * algorithms
 * Medium (58.13%)
 * Likes:    326
 * Dislikes: 0
 * Total Accepted:    65.6K
 * Total Submissions: 112.6K
 * Testcase Example:  '[[1,1,0],[1,1,0],[0,0,1]]'
 *
 * 班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是
 * C 的朋友。所谓的朋友圈，是指所有朋友的集合。
 *
 * 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j
 * 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * [[1,1,0],
 * ⁠[1,1,0],
 * ⁠[0,0,1]]
 * 输出：2
 * 解释：已知学生 0 和学生 1 互为朋友，他们在一个朋友圈。
 * 第2个学生自己在一个朋友圈。所以返回 2 。
 *
 *
 * 示例 2：
 *
 * 输入：
 * [[1,1,0],
 * ⁠[1,1,1],
 * ⁠[0,1,1]]
 * 输出：1
 * 解释：已知学生 0 和学生 1 互为朋友，学生 1 和学生 2 互为朋友，所以学生 0 和学生 2 也是朋友，所以他们三个在一个朋友圈，返回 1
 * 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= N <= 200
 * M[i][i] == 1
 * M[i][j] == M[j][i]
 *
 *
 */

// @lc code=start
func findCircleNum(M [][]int) int {
	return unionFind(M)
}

// 深度优先搜索
func methodWithDfs(M [][]int) int {
	var count int // 返回最终的结果
	visited := make([]int, len(M))  // visited[2]=1 代表某一行的第三个元素已经访问过了
	var dfs func (idx int)
	dfs = func (idx int)  {
		for j:=0; j<len(M); j++{
			if M[idx][j] == 1 && visited[j] == 0{
				visited[j] = 1
				dfs(j)
			}
		}
	}
	for i:=0; i<len(M); i++{
		if visited[i] == 0{
			dfs(i)
			count += 1
		}
	}

	return count
}


// 广度优先搜索
func methodWithBfs(M [][]int) int {
	var count int
	visited := make([]int, len(M))
	Q := make([]int, 0)
	for i:=0; i<len(M); i++{
		if visited[i] == 0{
			Q = append(Q, i)

			for len(Q) > 0{
				s := Q[0]
				Q = Q[1:]

				visited[s] = 1
				for j:=0; j<len(M); j++{
					if M[s][j] == 1 && visited[j] == 0{
						Q = append(Q, j)
					}
				}
			}

			count += 1
		}
	}

	return count
}


// 并查集的方法
func unionFind(M [][]int) int {
	uf := make([]int, len(M))
	for i:=0; i<len(M); i++{
		uf[i] = -1
	}

	for i:=0; i<len(M); i++{
		for j:=0; j<len(M[0]); j++{
			if M[i][j] == 1 && i!=j {
				set1 := find(i, uf)
				set2 := find(j, uf)
				if set1 != set2{
					uf[set1] = set2 
				}
			}
		}
	}

	fmt.Println(uf)
	var count int

	for _, v := range uf {
		if v == -1{
			count +=1
		}
 	}
	 return count
}
func find(n int, uf []int) int {
	if uf[n] == -1{
		return n
	}
	return find(uf[n], uf)

	// for uf[n] != -1{
	// 	n = uf[n]
	// }
	// return n
}


// @lc code=end
