/*
 * @lc app=leetcode.cn id=814 lang=golang
 *
 * [814] 二叉树剪枝
 *
 * https://leetcode-cn.com/problems/binary-tree-pruning/description/
 *
 * algorithms
 * Medium (72.87%)
 * Likes:    110
 * Dislikes: 0
 * Total Accepted:    12.2K
 * Total Submissions: 16.8K
 * Testcase Example:  '[1,null,0,0,1]'
 *
 * 给定二叉树根结点 root ，此外树的每个结点的值要么是 0，要么是 1。
 *
 * 返回移除了所有不包含 1 的子树的原二叉树。
 *
 * ( 节点 X 的子树为 X 本身，以及所有 X 的后代。)
 *
 *
 * 示例1:
 * 输入: [1,null,0,0,1]
 * 输出: [1,null,0,null,1]
 * ⁠
 * 解释:
 * 只有红色节点满足条件“所有不包含 1 的子树”。
 * 右图为返回的答案。
 *
 *
 *
 *
 *
 * 示例2:
 * 输入: [1,0,1,0,0,0,1]
 * 输出: [1,null,1,null,1]
 *
 *
 *
 *
 *
 *
 * 示例3:
 * 输入: [1,1,0,1,1,0,1,0]
 * 输出: [1,1,0,1,1,null,1]
 *
 *
 *
 *
 *
 * 说明:
 *
 *
 * 给定的二叉树最多有 100 个节点。
 * 每个节点的值只会为 0 或 1 。
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
func pruneTree(root *TreeNode) *TreeNode {
	return methodWithDfs(root)
}

func methodWithDfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// 先左右， 然后根节点
	root.Left = methodWithDfs(root.Left)
	root.Right = methodWithDfs(root.Right)

	// 如果二叉树的左节点、右节点都为空， 并且为0的话， 进行剪枝
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

// @lc code=end
