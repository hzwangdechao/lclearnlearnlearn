/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
 *
 * https://leetcode-cn.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (57.37%)
 * Likes:    394
 * Dislikes: 0
 * Total Accepted:    99.9K
 * Total Submissions: 171.9K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * 给定两个二叉树，编写一个函数来检验它们是否相同。
 *
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 *
 * 示例 1:
 *
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 *
 * ⁠       [1,2,3],   [1,2,3]
 *
 * 输出: true
 *
 * 示例 2:
 *
 * 输入:      1          1
 * ⁠         /           \
 * ⁠        2             2
 *
 * ⁠       [1,2],     [1,null,2]
 *
 * 输出: false
 *
 *
 * 示例 3:
 *
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 *
 * ⁠       [1,2,1],   [1,1,2]
 *
 * 输出: false
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// return recursive(p, q)
	return iteration(p, q)
}



func check(p *TreeNode, q *TreeNode) bool  {
	if p == nil && q == nil{
		return true
	}
	if p == nil || q == nil{
		return false
	}

	return p.Val == q.Val

}

func iteration(p *TreeNode, q *TreeNode) bool {
	// 创建队列
	queNode := []*TreeNode{p, q}
	for len(queNode) != 0{
		first := queNode[0]
		second := queNode[1]
		queNode = queNode[2:]

		x := check(first, second)
		fmt.Println(x)
		if !x{
			return false
		}

		if first != nil{
			queNode = append(queNode, first.Left)
			queNode = append(queNode, second.Left)
		}

		if second != nil {
			queNode = append(queNode, first.Right)
			queNode = append(queNode, second.Right)
		}

	}

	return true
}


func recursive(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil{
		return true
	}

	if p == nil || q == nil{
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
// @lc code=end
