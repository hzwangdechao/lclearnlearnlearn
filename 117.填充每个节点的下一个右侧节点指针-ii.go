/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 *
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (48.01%)
 * Likes:    166
 * Dislikes: 0
 * Total Accepted:    23.9K
 * Total Submissions: 48K
 * Testcase Example:  '[1,2,3,4,5,null,7]'
 *
 * 给定一个二叉树
 *
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 *
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 *
 * 初始状态下，所有 next 指针都被设置为 NULL。
 *
 *
 *
 * 进阶：
 *
 *
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 *
 *
 *
 *
 * 示例：
 *
 *
 *
 * 输入：root = [1,2,3,4,5,null,7]
 * 输出：[1,#,2,3,#,4,5,7,#]
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数小于 6000
 * -100 <= node.val <= 100
 *
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect2(root *Node) *Node {
	return bfs(root)


}
func connect(root *Node) *Node {
	if root == nil{
		return nil
	}

	start := root

	for start != nil{
		var last, nextStart *Node
		handle := func (cur *Node)  {
			if cur == nil{
				return
			}
			if nextStart == nil{
				nextStart = cur
			}

			if last != nil{
				last.Next = cur
			}
			last = cur
		}

		for p:=start; p!=nil; p=p.Next{
			handle(p.Left)
			handle(p.Right)
		}

		start = nextStart
	}
	return root
}
// 广度搜索+队列  与116题相似
func bfs(root *Node) *Node  {
	if root == nil{
		return nil
	}

	queNode := []*Node{root}
	for len(queNode) > 0{
		size := len(queNode)
		for i:=0; i<size; i++{
			node := queNode[0]
			queNode = queNode[1:]

			if i < size -1{
				node.Next = queNode[0]
			}

			if node.Left != nil{
				queNode = append(queNode, node.Left)
			}
			if node.Right != nil{
				queNode = append(queNode, node.Right)
			}

		}
	}
	return root
}
// @lc code=end
