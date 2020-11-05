/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 *
 * https://leetcode-cn.com/problems/redundant-connection/description/
 *
 * algorithms
 * Medium (60.42%)
 * Likes:    168
 * Dislikes: 0
 * Total Accepted:    18.9K
 * Total Submissions: 31.1K
 * Testcase Example:  '[[1,2],[1,3],[2,3]]'
 *
 * 在本问题中, 树指的是一个连通且无环的无向图。
 *
 * 输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N)
 * 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。
 *
 * 结果图是一个以边组成的二维数组。每一个边的元素是一对[u, v] ，满足 u < v，表示连接顶点u 和v的无向图的边。
 *
 * 返回一条可以删去的边，使得结果图是一个有着N个节点的树。如果有多个答案，则返回二维数组中最后出现的边。答案边 [u, v] 应满足相同的格式 u <
 * v。
 *
 * 示例 1：
 *
 * 输入: [[1,2], [1,3], [2,3]]
 * 输出: [2,3]
 * 解释: 给定的无向图为:
 * ⁠ 1
 * ⁠/ \
 * 2 - 3
 *
 *
 * 示例 2：
 *
 * 输入: [[1,2], [2,3], [3,4], [1,4], [1,5]]
 * 输出: [1,4]
 * 解释: 给定的无向图为:
 * 5 - 1 - 2
 * ⁠   |   |
 * ⁠   4 - 3
 *
 *
 * 注意:
 *
 *
 * 输入的二维数组大小在 3 到 1000。
 * 二维数组中的整数在1到N之间，其中N是输入数组的大小。
 *
 *
 * 更新(2017-09-26):
 * 我们已经重新检查了问题描述及测试用例，明确图是无向 图。对于有向图详见冗余连接II。对于造成任何不便，我们深感歉意。
 *
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	return union_find(edges)
}


// 并查集  参考地址  https://leetcode-cn.com/problems/redundant-connection/solution/tong-su-jiang-jie-bing-cha-ji-bang-zhu-xiao-bai-ku/
func union_find(edges [][]int) []int {
	rp := make([]int, 1001)
	size := len(edges)
	// 初始化个元素为单独的集合， 代表节点就是其本身
	for i:=0; i<size; i++{
		rp[i] = i
	}

	for i:=0; i<size; i++{
		// 找到边上两个节点所在集合的代表节点
		set1 := find(edges[i][0], rp)
		set2 := find(edges[i][1], rp)

		// 如果两个集合的代表节点相同，说明出现环， 返回打扮
		if set1 == set2{
			return edges[i]
		}else{
			// 两个集合独立， 合并集合。将前一个集合代表节点指到后一个集合代表节点上
			rp[set1] = set2
		}
	}
	return []int{0, 0}
}

// 查找路径并返回其代表节点
func find(n int, rp []int) int {
	num := n
	for rp[num] != num{
		num = rp[num]
	}
	return num
}


// @lc code=end
