/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (72.02%)
 * Likes:    295
 * Dislikes: 0
 * Total Accepted:    72.6K
 * Total Submissions: 100.7K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
 *
 * 说明：
 * 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
 *
 * 示例 1:
 *
 * 输入: root = [3,1,4,null,2], k = 1
 * ⁠  3
 * ⁠ / \
 * ⁠1   4
 * ⁠ \
 * 2
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: root = [5,3,6,2,4,null,null,1], k = 3
 * ⁠      5
 * ⁠     / \
 * ⁠    3   6
 * ⁠   / \
 * ⁠  2   4
 * ⁠ /
 * ⁠1
 * 输出: 3
 *
 * 进阶：
 * 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
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
func kthSmallest(root *TreeNode, k int) int {
	return methodWithBfs(root, k)
	// return methodWithDfs(root, k)
}

func methodWithBfs(root *TreeNode, k int) int {
	var stack []*TreeNode
	for{
		for root != nil{
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k -= 1
		if k == 0{
			return root.Val
		}
		root = root.Right
	}
}

func methodWithDfs(root *TreeNode, k int) int {
	var dfs func(node *TreeNode)[]int
	dfs = func (node *TreeNode) []int {
		if node == nil{
			return []int{}
		}

		res := dfs(node.Left)
		res = append(res, node.Val)
		res = append(res, dfs(node.Right)...)

		return res
	}

	return dfs(root)[k-1]
}

// @lc code=end
