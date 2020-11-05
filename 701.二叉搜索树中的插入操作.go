/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
 *
 * https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/description/
 *
 * algorithms
 * Medium (72.84%)
 * Likes:    71
 * Dislikes: 0
 * Total Accepted:    15.6K
 * Total Submissions: 21.4K
 * Testcase Example:  '[4,2,7,1,3]\n5'
 *
 * 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 保证原始二叉搜索树中不存在新值。
 *
 * 注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回任意有效的结果。
 *
 * 例如, 
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
 * 和 插入的值: 5
 *
 *
 * 你可以返回这个二叉搜索树:
 *
 *
 * ⁠        4
 * ⁠      /   \
 * ⁠     2     7
 * ⁠    / \   /
 * ⁠   1   3 5
 *
 *
 * 或者这个树也是有效的:
 *
 *
 * ⁠        5
 * ⁠      /   \
 * ⁠     2     7
 * ⁠    / \
 * ⁠   1   3
 * ⁠        \
 * ⁠         4
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {

	return iteration(root, val)


}
func iteration(root *TreeNode, val int) *TreeNode  {
	newNode := TreeNode{
		Val: val,
		Left: nil,
		Right: nil,
	}

	node := root
	for node != nil{
		if node.Val > val{
			// insert into left tree
			if node.Left == nil{
				node.Left = &newNode
				return root
			}else{
				node = node.Left
			}
		}else{
			// insert into right tree
			if node.Right == nil{
				node.Right = &newNode
				return root
			}else{
				node = node.Right
			}
		}
	}
	return &newNode
}


func recursive(root *TreeNode, val int) *TreeNode  {
	newNode := TreeNode{
		Val: val,
		Left: nil,
		Right: nil,
	}
	if root == nil{
		return &newNode
	}
	if root.Val > val{
		root.Left = recursive(root.Left, val)
	}else{
		root.Right = recursive(root.Right, val)
	}

	return root
}
// @lc code=end
