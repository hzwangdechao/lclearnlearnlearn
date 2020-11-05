/*
 * @lc app=leetcode.cn id=685 lang=golang
 *
 * [685] 冗余连接 II
 *
 * https://leetcode-cn.com/problems/redundant-connection-ii/description/
 *
 * algorithms
 * Hard (35.67%)
 * Likes:    141
 * Dislikes: 0
 * Total Accepted:    9.3K
 * Total Submissions: 22.3K
 * Testcase Example:  '[[1,2],[1,3],[2,3]]'
 *
 * 在本问题中，有根树指满足以下条件的有向图。该树只有一个根节点，所有其他节点都是该根节点的后继。每一个节点只有一个父节点，除了根节点没有父节点。
 *
 * 输入一个有向图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N)
 * 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。
 *
 * 结果图是一个以边组成的二维数组。 每一个边 的元素是一对 [u, v]，用以表示有向图中连接顶点 u 和顶点 v 的边，其中 u 是 v
 * 的一个父节点。
 *
 * 返回一条能删除的边，使得剩下的图是有N个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。
 *
 * 示例 1:
 *
 * 输入: [[1,2], [1,3], [2,3]]
 * 输出: [2,3]
 * 解释: 给定的有向图如下:
 * ⁠ 1
 * ⁠/ \
 * v   v
 * 2-->3
 *
 *
 * 示例 2:
 *
 * 输入: [[1,2], [2,3], [3,4], [4,1], [1,5]]
 * 输出: [4,1]
 * 解释: 给定的有向图如下:
 * 5 <- 1 -> 2
 * ⁠    ^    |
 * ⁠    |    v
 * ⁠    4 <- 3
 *
 *
 * 注意:
 *
 *
 * 二维数组大小的在3到1000范围内。
 * 二维数组中的每个整数在1到N之间，其中 N 是二维数组的大小。
 *
 *
 */

// @lc code=start
/*

*/
type unionFind struct{
	list []int
}

// 创建一个并查集
func newUnionFind(n int) unionFind {
	list := make([]int, n)
	for i:=0; i<n; i++{
		list[i] = i
	}
	return unionFind{list}
}

// 并查集的查询
func (uf unionFind)find(x int) int {
	// num := x
	// for uf.list[num] != num{
	// 	num = uf.list[num]
	// }
	// return num

	if uf.list[x] != x{
		uf.list[x] = uf.find(uf.list[x])
	}
	return uf.list[x]
}

// 并查集的合并
func (uf unionFind)union(from, to int)  {
	uf.list[uf.find(from)] = uf.find(to)
}


func findRedundantDirectedConnection(edges [][]int) []int {
	size := len(edges)
	uf := newUnionFind(size+1)
	parent := make([]int, size+1)  // parent[i] 表示i的父节点
	for i:=0; i<size+1; i++{
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int // 声明冲突边和环边
	for i:=0; i<size; i++{
		from, to := edges[i][0], edges[i][1]
		if parent[to] != to{
			// to有两个父节点
			conflictEdge = edges[i]
		}else{
			parent[to] = from  // 定义好to的父节点
			if uf.find(from) == uf.find(to) {
				// from 和 to已经连接
				cycleEdge = edges[i]
			}else{
				uf.union(from, to)
			}
		}
	}

	if conflictEdge == nil{
		return cycleEdge
	}

	if cycleEdge != nil{
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}



// @lc code=end
