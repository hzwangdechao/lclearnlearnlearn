/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
 *
 * https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/description/
 *
 * algorithms
 * Medium (66.26%)
 * Likes:    247
 * Dislikes: 0
 * Total Accepted:    58.9K
 * Total Submissions: 88.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
 *
 * 例如，从根到叶子节点路径 1->2->3 代表数字 123。
 *
 * 计算从根到叶子节点生成的所有数字之和。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * 输出: 25
 * 解释:
 * 从根到叶子节点路径 1->2 代表数字 12.
 * 从根到叶子节点路径 1->3 代表数字 13.
 * 因此，数字总和 = 12 + 13 = 25.
 *
 * 示例 2:
 *
 * 输入: [4,9,0,5,1]
 * ⁠   4
 * ⁠  / \
 * ⁠ 9   0
 * / \
 * 5   1
 * 输出: 1026
 * 解释:
 * 从根到叶子节点路径 4->9->5 代表数字 495.
 * 从根到叶子节点路径 4->9->1 代表数字 491.
 * 从根到叶子节点路径 4->0 代表数字 40.
 * 因此，数字总和 = 495 + 491 + 40 = 1026.
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

func sumNumbersmethodWithBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	nodeQ := []*TreeNode{root}
	valueQ := []int{root.Val}

	for len(nodeQ) > 0 {
		node := nodeQ[0]
		nodeQ = nodeQ[1:]
		value := valueQ[0]
		valueQ = valueQ[1:]

		if node.Left == nil && node.Right == nil {
			res += value
		}

		if node.Left != nil {
			nodeQ = append(nodeQ, node.Left)
			valueQ = append(valueQ, value*10+node.Left.Val)
		}

		if node.Right != nil {
			nodeQ = append(nodeQ, node.Right)
			valueQ = append(valueQ, value*10+node.Right.Val)
		}
	}

	return res
}

func sumNumbers(root *TreeNode) int {
	var (
		dfs func(node *TreeNode, preSum int) int
	)
	dfs = func(node *TreeNode, preSum int) int {
		if node == nil {
			return 0
		}

		sum := preSum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}

		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}

	return dfs(root, 0)
}

// @lc code=end
