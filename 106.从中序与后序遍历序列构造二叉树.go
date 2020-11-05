/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (68.16%)
 * Likes:    234
 * Dislikes: 0
 * Total Accepted:    41.3K
 * Total Submissions: 59.9K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// return method1(inorder, postorder)

	l := len(postorder)

	for i, v := range inorder{
		if v == postorder[l-1]{
			return &TreeNode{
				Val: v,
				Left: buildTree(inorder[:i], postorder[:i]),
				Right: buildTree(inorder[i+1:], postorder[i:l-1]),
			}
		}
	}

	return nil

}




func buildTree(inorder []int, postorder []int) *TreeNode {

	l := len(postorder)
	for i, v := range inorder{
		if v == postorder[l-1]{
			return &TreeNode{
				Val: v,
				Left: buildTree(inorder[:i], postorder[:i]),
				Right:buildTree(inorder[i+1:], postorder[i:l-1]),
			}
		}
	}
}


// @lc code=end
