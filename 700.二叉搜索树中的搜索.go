/*
 * @lc app=leetcode.cn id=700 lang=golang
 *
 * [700] 二叉搜索树中的搜索
 *
 * https://leetcode-cn.com/problems/search-in-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (74.06%)
 * Likes:    72
 * Dislikes: 0
 * Total Accepted:    29.5K
 * Total Submissions: 39.9K
 * Testcase Example:  '[4,2,7,1,3]\n2'
 *
 * 给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
 *
 * 例如，
 *
 *
 * 给定二叉搜索树:
 *
 * ⁠       4
 * ⁠      / \
 * ⁠     2   7
 * ⁠    / \
 * ⁠   1   3
 *
 * 和值: 2
 *
 *
 * 你应该返回如下子树:
 *
 *
 * ⁠     2
 * ⁠    / \
 * ⁠   1   3
 *
 *
 * 在上述示例中，如果要找的值是 5，但因为没有节点值为 5，我们应该返回 NULL。
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
func searchBST(root *TreeNode, val int) *TreeNode {
	return iteration(root, val)
}

// 迭代法
func iteration(root *TreeNode, val int) *TreeNode   {
	if root == nil {
		return root
	}
	for root != nil{
		if root.Val == val{
			return root
		}else if root.Val > val{
			root = root.Left
		}else{
			root = root.Right
		}
	}
	return nil
}

// 递归法
func recursive(root *TreeNode, val int) *TreeNode   {
	if root == nil || root.Val == val{
		return root
	}

	if root.Val > val{
		return recursive(root.Left, val)
	}else{
		return recursive(root.Right, val)
	}
}
// @lc code=end
