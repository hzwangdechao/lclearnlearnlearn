/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 *
 * https://leetcode-cn.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (49.60%)
 * Likes:    363
 * Dislikes: 0
 * Total Accepted:    107.4K
 * Total Submissions: 212.3K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例: 
 * 给定如下二叉树，以及目标和 sum = 22，
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \      \
 * ⁠       7    2      1
 *
 *
 * 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
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
func hasPathSum(root *TreeNode, sum int) bool {
	// return recursive(root, sum)
	return queue(root, sum)
}

// 利用节点队列和值队列求解
func queue(root *TreeNode, sum int) bool  {
	if root == nil{
		return false
	}
	queNode := []*TreeNode{}
	queVal := []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) != 0{
		curNode := queNode[0]
		queNode = queNode[1:]
		curVal := queVal[0]
		queVal = queVal[1:]

		if curNode.Left == nil && curNode.Right == nil{
			if curVal == sum{
				return true
			}
			continue
		}
		if curNode.Left != nil{
			queNode = append(queNode, curNode.Left)
			queVal = append(queVal, curNode.Left.Val + curVal)
		}
		if curNode.Right != nil{
			queNode = append(queNode, curNode.Right)
			queVal = append(queVal, curNode.Right.Val + curVal)
		}
	}
	return false

}

// 递归法求解
func recursive(root *TreeNode, sum int) bool  {
	if root == nil{
		return false
	}

	if root.Left == nil && root.Right == nil{
		return root.Val == sum
	}

	return recursive(root.Right, sum-root.Val) || recursive(root.Left, sum-root.Val)

}
// @lc code=end
