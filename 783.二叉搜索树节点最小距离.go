/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
 *
 * https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (53.99%)
 * Likes:    88
 * Dislikes: 0
 * Total Accepted:    19.3K
 * Total Submissions: 35.4K
 * Testcase Example:  '[4,2,6,1,3,null,null]'
 *
 * 给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。
 *
 *
 *
 * 示例：
 *
 * 输入: root = [4,2,6,1,3,null,null]
 * 输出: 1
 * 解释:
 * 注意，root是树节点对象(TreeNode object)，而不是数组。
 *
 * 给定的树 [4,2,6,1,3,null,null] 可表示为下图:
 *
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \
 * ⁠   1   3
 *
 * 最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
 *
 *
 *
 * 注意：
 *
 *
 * 二叉树的大小范围在 2 到 100。
 * 二叉树总是有效的，每个节点的值都是整数，且不重复。
 * 本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
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
func minDiffInBST(root *TreeNode) int {
	return method2(root)
}

// 利用中序遍历的方法, 计算相邻节点之间的差值
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

// 将二叉搜索树的节点存放到一个数组中， 然后对数组进行排序， 然后获取相邻节点查的最小值
func method1(root *TreeNode) int {
	var (
		vals []int
		res  = math.MaxInt32
		dfs  func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node != nil {
			vals = append(vals, node.Val)
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	// 对vals进行排序
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
