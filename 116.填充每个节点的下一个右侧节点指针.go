/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 *
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/description/
 *
 * algorithms
 * Medium (58.47%)
 * Likes:    197
 * Dislikes: 0
 * Total Accepted:    38.5K
 * Total Submissions: 63.4K
 * Testcase Example:  '[1,2,3,4,5,6,7]'
 *
 * 给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
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
 * 示例：
 *
 *
 *
 *
 * 输入：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":{"$id":"6","left":null,"next":null,"right":null,"val":6},"next":null,"right":{"$id":"7","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}
 *
 *
 * 输出：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":{"$id":"4","left":null,"next":{"$id":"5","left":null,"next":{"$id":"6","left":null,"next":null,"right":null,"val":7},"right":null,"val":6},"right":null,"val":5},"right":null,"val":4},"next":{"$id":"7","left":{"$ref":"5"},"next":null,"right":{"$ref":"6"},"val":3},"right":{"$ref":"4"},"val":2},"next":null,"right":{"$ref":"7"},"val":1}
 *
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
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

func connect(root *Node) *Node {
	// return dfs2(root)
	// return bfs1(root)

	if root == nil{
		return root
	}
	connectTwoNode(root.Left, root.Right)

	return root
}

func bfs1(root *Node)*Node  {
	if root == nil{
		return nil
	}
	// 构建队列
	queNode := []*Node{root}
	for len(queNode) > 0{
		size := len(queNode)
		for i := 0; i < size; i++{
			fmt.Println(i, size)
			// 从队列中取出节点
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

func dfs2(root *Node) *Node {
	if root == nil{
		return nil
	}
	leftmost := root
	for leftmost.Left != nil{
		node := leftmost
		for node != nil{
			node.Left.Next = node.Right
			if node.Next != nil{
				node.Right.Next = node.Next.Left
			}
			node = node.Next
		}
		leftmost = leftmost.Left
	}
	return root
}

func dfs1(root *Node) *Node {
	if root == nil{
		return nil
	}
	l := root.Left
	r := root.Right
	for l != nil{
		l.Next = r
		l = l.Right
		r = r.Left
	}

	dfs1(root.Left)
	dfs1(root.Right)

	return root
}


func connectTwoNode(node1, node2 *Node)  {
	if node1 == nil || node2 == nil{
		return
	}
	node1.Next = node2
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)
	connectTwoNode(node1.Right, node2.Left)
}
// @lc code=end
