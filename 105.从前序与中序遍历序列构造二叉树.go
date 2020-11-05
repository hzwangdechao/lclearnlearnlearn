/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (65.12%)
 * Likes:    553
 * Dislikes: 0
 * Total Accepted:    92.2K
 * Total Submissions: 136.9K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	return recursive(preorder, inorder)

}




func recursive(preorder []int, inorder []int) *TreeNode  {
	// 遍历 中序列表， 与前序遍历的第一个元素进行比较
	for i, v := range inorder{
		if v == preorder[0]{
			return &TreeNode{
				Val: v,
				Left: recursive(preorder[1:i+1], inorder[0:i]),  // 中序列表和前序列表的长度应该是一样的
				Right: recursive(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}
// @lc code=end
