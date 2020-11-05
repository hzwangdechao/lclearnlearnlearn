/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (65.46%)
 * Likes:    280
 * Dislikes: 0
 * Total Accepted:    116.3K
 * Total Submissions: 176.7K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
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
 * 输出: [1,2,3]
 *
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
func preorderTraversal(root *TreeNode) []int {
	// return preorderRecursive(root)
	// return preorderIterate(root)
	// return bfs(root)
	return dfs(root)
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
			if cur.Left != nil{
				stack.Push(cur.Left)
			}
			stack.Push(cur)
			stack.Push(nil)
		}else{
			res = append(res, stack.Pop().Val)
		}
	}
	return res

}



// 递归方法  根   左  右
func preorderRecursive(root *TreeNode)[]int  {
	if root == nil{
		return []int{}
	}

	res := append([]int{root.Val}, preorderRecursive(root.Left)...)
	res = append(res, preorderRecursive(root.Right)...)
	return res
}

// 广度优先实现二叉树的前序遍历
func bfs(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}

	stack := []*TreeNode{root}
	res := []int{}

	for len(stack)> 0{
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil{
			if cur.Right != nil{
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil{
				stack = append(stack,  cur.Left)
			}

			stack = append(stack, cur)
			stack = append(stack, nil)
		}else{
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, last.Val)
		}
	}

	return res
}

// 深度优先搜索实现二叉树的前序遍历  根   左   右
func dfs(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}

	res := append([]int{root.Val}, dfs(root.Left)...)
	res = append(res, dfs(root.Right)...)
	return res
}

// @lc code=end
