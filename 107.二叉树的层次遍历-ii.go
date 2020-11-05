/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (66.22%)
 * Likes:    311
 * Dislikes: 0
 * Total Accepted:    86.8K
 * Total Submissions: 129.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其自底向上的层次遍历为：
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	res := make([][]int, 0)
	Q := []*TreeNode{root}

	for i:=0; len(Q)>0; i++{
		nextQueue := make([]*TreeNode, 0)
		res = append(res, []int{})

		for j:=0; j<len(Q); j++{
			node := Q[j]
			res[i] = append(res[i], node.Val)

			if node.Left != nil{
				nextQueue = append(nextQueue, node.Left)
			}

			if node.Right != nil{
				nextQueue = append(nextQueue, node.Right)
			}
		}
		Q = nextQueue

	}

	// 对res中的数据进行翻转
	for i:=0; i<len(res)/2; i++{
		res[i], res[len(res)-i-1] =  res[len(res)-i-1], res[i]
	}
	return res
}
// @lc code=end
