/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (71.30%)
 * Likes:    532
 * Dislikes: 0
 * Total Accepted:    173.2K
 * Total Submissions: 241.4K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [1,3,2]
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
func inorderTraversal(root *TreeNode) []int {
	// return preorderRecursive(root)
	return preorderIterate(root)
}

// 迭代法
type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}
func preorderIterate(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	stack, res := Stack([]*TreeNode{root}), []int{}
	for len(stack) > 0{
		cur := stack.Pop()
		if cur != nil{

			if cur.Right != nil{
				stack.Push(cur.Right)
			}

			stack.Push(cur)
			stack.Push(nil)

			if cur.Left != nil{
				stack.Push(cur.Left)
			}

		}else{
			res = append(res, stack.Pop().Val)
		}
	}

	return res
}

// 递归法  左  根  右  
func preorderRecursive(root *TreeNode)[]int  {
	if root == nil{
		return []int{}
	}

	res := append(preorderRecursive(root.Left), root.Val)
	res = append(res, preorderRecursive(root.Right)...)
	return res
}

func bfs(root *TreeNode)[]int  {
	if root == nil{
		return []int{}
	}
	stack := []*TreeNode{root}
	res := []int{}

	for len(stack) > 0{
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur != nil{
			if cur.Right != nil{
				stack = append(stack, cur.Right)
			}
			stack = append(stack, cur)
			stack = append(stack, nil)

			if cur.Left != nil{
				stack = append(stack, cur.Left)
			}
		}else{
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, last.Val)
		}

	}

	return res
}
// @lc code=end
