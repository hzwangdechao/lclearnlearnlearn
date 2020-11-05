/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 *
 * https://leetcode-cn.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (60.28%)
 * Likes:    402
 * Dislikes: 0
 * Total Accepted:    89.9K
 * Total Submissions: 149.1K
 * Testcase Example:  '19'
 *
 * 编写一个算法来判断一个数 n 是不是快乐数。
 *
 * 「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到
 * 1。如果 可以变为  1，那么这个数就是快乐数。
 *
 * 如果 n 是快乐数就返回 True ；不是，则返回 False 。
 *
 *
 *
 * 示例：
 *
 * 输入：19
 * 输出：true
 * 解释：
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 *
 */

// @lc code=start
func isHappy(n int) bool {
	m := make(map[int]bool, 0)
	for n!= 1{
		// 如果n不等于1就一直遍历
		n = calc(n)
		if _, ok := m[n]; !ok{
			// 不存在循环的话就将num放到m中
			m[n] = true
		}else{
			// 存在的话说明出现了循环
			return false
		}
	}
	return true
}

// 计算整数的平方和
func calc(n int)int  {
	res := 0
	for n != 0{
		res += (n%10)*(n%10)
		n /= 10
	}
	return res
}
// @lc code=end
