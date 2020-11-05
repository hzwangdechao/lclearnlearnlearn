/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (60.98%)
 * Likes:    344
 * Dislikes: 0
 * Total Accepted:    87.7K
 * Total Submissions: 143.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 *
 *
 * 返回:
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
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
func pathSum(root *TreeNode, sum int) [][]int {
	return methodWithDfs(root, sum)
}

func methodWithDfs(root *TreeNode, sum int) [][]int {
	var res [][]int

	var dfs func(*TreeNode, int, []int)
	dfs = func (node *TreeNode, target int, path []int)  {
		if node == nil{
			return
		}

		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil{
			if target == node.Val{
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)

			}
		}

		dfs(node.Left, target-node.Val, path)
		dfs(node.Right, target-node.Val, path)
	}

	dfs(root, sum, []int{})

	return res
}
// @lc code=end
