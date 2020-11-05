/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (61.71%)
 * Likes:    536
 * Dislikes: 0
 * Total Accepted:    146.9K
 * Total Submissions: 233.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层次遍历结果：
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
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
// func levelOrder(root *TreeNode) [][]int {
// 	res :=[][]int{}  // 用于返回结果
// 	if root == nil{
// 		return res
// 	}

// 	// 模拟实现队列
// 	q := []*TreeNode{root}
// 	for i := 0; len(q) > 0; i++{
// 		res = append(res, []int{})
// 		p := []*TreeNode{}
// 		for j := 0; j < len(q); j++{
// 			node := q[j]
// 			res[i] = append(res[i], node.Val)
// 			if node.Left != nil{
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil{
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return res
// }

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}  // 存放最终返回的结果
	if root == nil{
		return res
	}

	Q := []*TreeNode{root}  // 当前层的节点队列
	for i := 0; len(Q) > 0; i ++{
		NQ := []*TreeNode{}  // 下一层的节点队列
		res = append(res, []int{})

		for j := 0; j < len(Q); j ++{
			// 遍历当前层的节点队列
			node := Q[j] // 取出一个节点
			res[i] = append(res[i], node.Val)

			// 如果当前节点有子节点的话， 则将其添加到NQ中
			if node.Left != nil{
				NQ = append(NQ, node.Left)
			}
			if node.Right != nil{
				NQ = append(NQ, node.Right)
			}
		}
		// 将NQ(下一层节点)作为接下来要遍历的队列
		Q = NQ
	}
	return res
}

// @lc code=end
