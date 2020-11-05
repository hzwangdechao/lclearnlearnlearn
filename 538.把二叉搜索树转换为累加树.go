/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 *
 * https://leetcode-cn.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (63.32%)
 * Likes:    356
 * Dislikes: 0
 * Total Accepted:    50.5K
 * Total Submissions: 78.9K
 * Testcase Example:  '[5,2,13]'
 *
 * 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater
 * Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
 *
 *
 *
 * 例如：
 *
 * 输入: 原始二叉搜索树:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * 输出: 转换为累加树:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
 *
 *
 *
 *
 * 注意：本题和 1038:
 * https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同
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
func convertBST(root *TreeNode) *TreeNode {
	return convertBST2(root)
}


func convertBST1(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func (node *TreeNode)  {
		if node == nil{
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	node := root

	for node != nil{
		if node.Right == nil{
			// 如果当前节点的右节点为空的话, 处理当前节点， 并继续遍历当前节点的左子节点
			sum += node.Val
			node.Val = sum
			node = node.Left
		}else{
			// 如果当前节点的右节点不为空的话， 找到当前节点右子树的最左节点
			succ := getSuccessor(node)
			if succ.Left == nil{
				// 如果最左节点的左指针为空， 将最左节点的左指针指向当前节点， 遍历当前节点的右子节点
				succ.Left = node
				node = node.Right
			}else{
				/// 将最左节点的左指针重置为空， 处理当前节点， 并将当前节点重置为其左节点
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}
func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node{
		succ = succ.Left
	}
	return succ
}



// @lc code=end
