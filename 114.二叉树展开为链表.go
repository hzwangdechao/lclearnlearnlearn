/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (69.62%)
 * Likes:    448
 * Dislikes: 0
 * Total Accepted:    56.9K
 * Total Submissions: 81.2K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给定一个二叉树，原地将它展开为一个单链表。
 *
 *
 *
 * 例如，给定二叉树
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 *
 * 将其展开为：
 *
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
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


// 前序遍历+展开 同时进行
func flatten222(root *TreeNode)  {
    if root == nil {
        return
    }
    stack := []*TreeNode{root}
    var prev *TreeNode
    for len(stack) > 0 {
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if prev != nil {
            prev.Left, prev.Right = nil, curr
        }
        left, right := curr.Left, curr.Right
        if right != nil {
            stack = append(stack, right)
        }
        if left != nil {
            stack = append(stack, left)
        }
        prev = curr
    }
}

func flatten(root *TreeNode)  {
	if root == nil{
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left
	p := root
	for p.Right!=nil{
		p = p.Right
	}
	p.Right = right
}



//前序遍历+迭代
func flatten2(root *TreeNode)  {
	list := []*TreeNode{}
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0{
		// 将二叉树的左节点放到list和stack照片那个
		for node != nil{
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
        node = node.Right
        stack = stack[:len(stack)-1]

	}
	for i := 1; i < len(list); i++ {
        prev, curr := list[i-1], list[i]
        prev.Left, prev.Right = nil, curr
    }
}

 //  前序遍历+递归
func flatten1(root *TreeNode)  {
	list := preorderTraversal(root)
	for i:=1; i<len(list); i++{
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}

}

//前序遍历+递归
func preorderTraversal(root *TreeNode)[]*TreeNode  {
	list := []*TreeNode{}
	if root != nil{
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}
// @lc code=end
