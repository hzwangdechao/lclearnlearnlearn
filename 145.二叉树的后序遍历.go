/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (71.26%)
 * Likes:    328
 * Dislikes: 0
 * Total Accepted:    85.9K
 * Total Submissions: 119.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
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
 * 输出: [3,2,1]
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
func postorderTraversal(root *TreeNode) []int {
	// return preorderRecursive(root)
	// return preorderIterate(root)
	// return dfs(root)
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
	if root  ==  nil{
		return []int{}
	}
		stack, res := Stack([]*TreeNode{root}), []int{}
		for len(stack) > 0{
			cur := stack.Pop()
			if cur != nil{
				stack.Push(cur)
				stack.Push(nil)

				if cur.Right != nil{
					stack.Push(cur.Right)
				}

				if cur.Left != nil{
					stack.Push(cur.Left)
				}
			}else{
				res = append(res, stack.Pop().Val)
			}
		}
	return res

}
// 递归法
func preorderRecursive(root *TreeNode)[]int  {
	if root == nil{
		return []int{}
	}

	res := append(preorderRecursive(root.Left), preorderRecursive(root.Right)...)
	res = append(res, root.Val)
	return res

}

// 深度优先搜索实现二叉树的后序遍历      左   右   根
func dfs(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}

	res := append(dfs(root.Left), dfs(root.Right)...)
	res =  append(res, root.Val)
	return res
}

// 广度优先搜索实现二叉树的后序遍历   左   右   根
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
			stack = append(stack, cur)
			stack = append(stack, nil)

			if cur.Right != nil{
				stack = append(stack, cur.Right)
			}

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
