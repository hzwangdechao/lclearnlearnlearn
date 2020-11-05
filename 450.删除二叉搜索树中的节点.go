/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
 *
 * https://leetcode-cn.com/problems/delete-node-in-a-bst/description/
 *
 * algorithms
 * Medium (42.38%)
 * Likes:    238
 * Dislikes: 0
 * Total Accepted:    17.1K
 * Total Submissions: 40.4K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n3'
 *
 * 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key
 * 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
 *
 * 一般来说，删除节点可分为两个步骤：
 *
 *
 * 首先找到需要删除的节点；
 * 如果找到了，删除它。
 *
 *
 * 说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
 *
 * 示例:
 *
 *
 * root = [5,3,6,2,4,null,7]
 * key = 3
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * 给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
 *
 * 一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 4   6
 * ⁠/     \
 * 2       7
 *
 * 另一个正确答案是 [5,2,6,null,4,null,7]。
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 2   6
 * ⁠  \   \
 * ⁠   4   7
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

 /*
 如果当前节点的值比要删除的元素大的话，迭代删除当前节点的左子树， 比要删除的元素的值小的话， 迭代删除当前节点的右子树
 相对的情况下  如果当前节点没有子节点， 则直接删除
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil{
		return root
	}

	if root.Val > key{
		root.Left = deleteNode(root.Left, key)
	}else if root.Val < key{
		root.Right = deleteNode(root.Right, key)
	}else{
		if root.Left == nil{
			return root.Right
		}
		if root.Right ==  nil{
			return root.Left
		}

		// 找出左子树中的最大节点 或者 找出右子树中的最小节点 来替代
		maxNode := getMaxNode(root.Left)
		root.Val = maxNode.Val
		root.Left = deleteNode(root.Left, maxNode.Val)
	}

	return root
}


// 找出最大的元素
func getMaxNode(root *TreeNode) *TreeNode {
	for root.Right != nil{
		root = root.Right
	}
	return root
}

// 找出最小的节点
func getMinNode(root *TreeNode)*TreeNode  {
	for root.Left != nil{
		root = root.Left
	}
	return root
}

// @lc code=end
