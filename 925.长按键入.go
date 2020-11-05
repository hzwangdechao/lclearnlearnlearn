/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 *
 * https://leetcode-cn.com/problems/long-pressed-name/description/
 *
 * algorithms
 * Easy (39.82%)
 * Likes:    101
 * Dislikes: 0
 * Total Accepted:    15.9K
 * Total Submissions: 39.7K
 * Testcase Example:  '"alex"\n"aaleex"'
 *
 * 你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
 *
 * 你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。
 *
 *
 *
 * 示例 1：
 *
 * 输入：name = "alex", typed = "aaleex"
 * 输出：true
 * 解释：'alex' 中的 'a' 和 'e' 被长按。
 *
 *
 * 示例 2：
 *
 * 输入：name = "saeed", typed = "ssaaedd"
 * 输出：false
 * 解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。
 *
 *
 * 示例 3：
 *
 * 输入：name = "leelee", typed = "lleeelee"
 * 输出：true
 *
 *
 * 示例 4：
 *
 * 输入：name = "laiden", typed = "laiden"
 * 输出：true
 * 解释：长按名字中的字符并不是必要的。
 *
 *
 *
 *
 * 提示：
 *
 *
 * name.length <= 1000
 * typed.length <= 1000
 * name 和 typed 的字符都是小写字母。
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	// 循环遍历typed
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] { // 防止越界， 并且两个字符相等的情况
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] { // 两个字符不相等的情况， 如果typed当前字符和上一个字符相等的话， j++
			j++
		} else {
			return false
		}
	}

	return i == len(name)
}
func isLongPressedNameWdc(name string, typed string) bool {
	// 键入的名字比实际名字短的话， 直接返回false
	if len(name) > len(typed) {
		return false
	}

	j := 0                       // 从头开始遍历 typed
	for i := 0; i < len(name); { // 循环遍历 name 字符串
		if j >= len(typed) { // pyplrz  ppyypllr  typed已经循环遍历完但是name还没遍历完的情况
			return false
		}

		if name[i] != typed[j] { // name 和 typed 对应位置不相等的情况
			return false
		}

		//  lleecde  llleeecddee  typed中连续出现的次数必须比name中连续出现的次数大
		nameCnt := 0
		for i < len(name)-1 && name[i] == name[i+1] {
			nameCnt++
			i++
		}
		typedCnt := 0
		for j < len(typed)-1 && typed[j+1] == typed[j] {
			typedCnt++
			j++
		}
		if typedCnt < nameCnt {
			return false
		}

		// 移动两个字符串中的指针
		i++
		j++
	}

	// name已经检查完毕， 但typed中有不在name中的字符
	if j < len(typed) { // alex  alexxr
		return false
	}

	// 没有出现错误返回true
	return true
}

// @lc code=end
