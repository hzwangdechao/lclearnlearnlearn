/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 *
 * https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (58.13%)
 * Likes:    146
 * Dislikes: 0
 * Total Accepted:    25.7K
 * Total Submissions: 43.4K
 * Testcase Example:  '[1,null,3,2]'
 *
 * 给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。
 *
 *
 *
 * 示例：
 *
 * 输入：
 *
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 *
 * 输出：
 * 1
 *
 * 解释：
 * 最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中至少有 2 个节点。
 * 本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/
 * 相同
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	return method2(root)
}

// 利用中序遍历的方法， 计算相邻元素的差值然后取最小值
func method2(root *TreeNode) int {
	var (
		res  = math.MaxInt32
		prev = math.MinInt32
		dfs  func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			res = min(res, node.Val-prev)
			prev = node.Val
			dfs(node.Right)
		}
	}
	dfs(root)
	return res
}

// 将二叉搜索中的节点放到一个数组中，然后对数组进行排序， 遍历数组找到做小的差值
func method1(root *TreeNode) int {
	var (
		vals []int
		dfs  func(node *TreeNode)
		res  = math.MaxInt32
	)
	dfs = func(node *TreeNode) {
		if node != nil {
			vals = append(vals, node.Val)
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	// 对数组进行排序
	sort.Ints(vals)

	for i := 1; i < len(vals); i++ {
		res = min(res, vals[i]-vals[i-1])
	}

	return res

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
