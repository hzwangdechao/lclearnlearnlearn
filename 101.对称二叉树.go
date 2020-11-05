/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (50.89%)
 * Likes:    861
 * Dislikes: 0
 * Total Accepted:    162.6K
 * Total Submissions: 311.4K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 *
 *
 *
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 *
 *
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 *
 *
 * 进阶：
 *
 * 你可以运用递归和迭代两种方法解决这个问题吗？
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
func isSymmetric(root *TreeNode) bool {
    // return recursive(root, root)
    return iterate(root)

}

// 迭代法 + 队列实现
func iterate(root *TreeNode) bool{
    root1, root2 := root, root

    // 模拟实现队列, 一次性放两个节点， 一次性读取两个节点
    q := []*TreeNode{}
    q = append(q, root1)
    q = append(q, root2)

    for len(q) > 0{
        root1, root2 = q[0], q[1]
        q = q[2:]

        if root1 == nil && root2 == nil{
            continue
        }
        if root1 == nil || root2 == nil{
            return false
        }

        if root1.Val != root2.Val{
            return false
        }

        q = append(q, root1.Left)
        q = append(q, root2.Right)

        q = append(q, root1.Right)
        q = append(q, root2.Left)

    }
    return true

}

// 递归法遍历
func recursive(root1, root2 *TreeNode) bool{
    // 如果两棵树都是空的话， 返回true
    if root1 == nil && root2 == nil{
        return true
    }
    // 如果只有一棵为空的话， 返回false
    if root1 == nil || root2 == nil{
        return false
    }

    return root1.Val == root2.Val  && recursive(root1.Left, root2.Right) && recursive(root1.Right, root2.Left)
}
// @lc code=end
