/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (42.23%)
 * Likes:    271
 * Dislikes: 0
 * Total Accepted:    82.4K
 * Total Submissions: 193.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最小深度  2.
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



 // 使用广度遍历二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	res := 1
	// 构建队列， 存放二叉树每一层上面的节点
	q := []*TreeNode{root}

	for len(q) > 0{
		x := len(q)
		// 将当队列中的所有节点向四周扩散
		for i:=0; i<x; i++{
			cur := q[0]
			q = q[1:]
			// 判断当前节点是否有左右节点
			if cur.Left == nil && cur.Right == nil{
				return res
			}

			if cur.Left != nil{
				q = append(q, cur.Left)
			}
			if cur.Right != nil{
				q = append(q, cur.Right)
			}
		}
		res += 1
	}

	return res

}







 // 使用深度遍历二叉树的做小深度
func minDepthDfs(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if root.Left == nil && root.Right == nil{
		return 1
	}
	if root.Left != nil && root.Right == nil{
		return minDepth(root.Left) + 1
	}
	if root.Right != nil && root.Left == nil{
		return minDepth(root.Right) + 1
	}

	return min(minDepth(root.Right), minDepth(root.Left)) + 1
}

func min(a , b int) int  {
	if a < b{
		return a
	}
	return b

}

// @lc code=end
