/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (30.61%)
 * Likes:    622
 * Dislikes: 0
 * Total Accepted:    129.6K
 * Total Submissions: 412.5K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 *
 * 假设一个二叉搜索树具有如下特征：
 *
 *
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 * 示例 1:
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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



// func isValidBST(root *TreeNode) bool {
// 	// return recurse(root, math.MinInt64, math.MaxInt64)
// 	return inorderTraversal(root)

// }

// 迭代法中序遍历
type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}
func inorderTraversal(root *TreeNode) bool  {
	if root == nil{
		return true
	}
	stack := Stack([]*TreeNode{root})
	inorder := math.MinInt64
	for len(stack) > 0{
		cur := stack.Pop()
		fmt.Println(cur)
		if cur != nil{
			if cur.Right != nil{
				stack.Push(cur.Right)
			}
			stack.Push(cur)
			stack.Push(nil)  // 终止遍历
			if cur.Left != nil{
				stack.Push(cur.Left)
			}
		}else{
			val := stack.Pop().Val
			if val <= inorder{
				return false
			}
			inorder =  val
		}
	}
	return true

}

// 递归法判断二叉搜索树
func recurse(node *TreeNode, lower int, upper int) bool {
	if node == nil{
		return true
	}
	val := node.Val

	// 不能小于下界, 不能大于上界
	if val <= lower || val >= upper{
		return false
	}

	// 将当前值作为上界， 继续比较左子树; 将当前值作为下界， 继续比较右子树

	return recurse(node.Left, lower, val) && recurse(node.Right, val, upper)

}



func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
	inorder := math.MinInt64
	fmt.Println(len(stack), root)
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
			root = root.Left
		}
		fmt.Println("done")
		fmt.Println(stack)
		root = stack[len(stack)-1]
		fmt.Println(root)
		fmt.Println(root.Right)

        stack = stack[:len(stack)-1]
        if root.Val <= inorder {
            return false
        }
        inorder = root.Val
        root = root.Right
    }
    return true
}

// @lc code=end
