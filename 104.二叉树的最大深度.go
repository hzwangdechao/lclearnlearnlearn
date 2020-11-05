/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (73.05%)
 * Likes:    566
 * Dislikes: 0
 * Total Accepted:    189.5K
 * Total Submissions: 258K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
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
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	// lLehgth := maxDepth(root.Left) // 左子树的最大深度
	// rLength := maxDepth(root.Right)  // 右子树的最大深度

	// return max(lLehgth, rLength) + 1

	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil && root.Right == nil{
		return maxDepth(root.Left) + 1
	}
	if root.Left == nil && root.Right != nil {
		return maxDepth(root.Right) + 1
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b

}
// @lc code=end
