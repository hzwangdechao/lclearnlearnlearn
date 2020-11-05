/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 *
 * https://leetcode-cn.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (55.06%)
 * Likes:    461
 * Dislikes: 0
 * Total Accepted:    93.5K
 * Total Submissions: 169.6K
 * Testcase Example:  '1'
 *
 * 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 被读作  "one 1"  ("一个一") , 即 11。
 * 11 被读作 "two 1s" ("两个一"）, 即 21。
 * 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
 *
 * 给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
 *
 * 注意：整数序列中的每一项将表示为一个字符串。
 *
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "1"
 * 解释：这是一个基本样例。
 *
 * 示例 2:
 *
 * 输入: 4
 * 输出: "1211"
 * 解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 =
 * 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
 *
 */

// @lc code=start
// func countAndSay(n int) string {
// 	// 直观解法， 先得到前一项， 再计算后一项
// 	if n ==1 {
// 		return "1"
// 	}
// 	prevStr := "1" // 上一个
// 	curChar := "1" // 当前
// 	curCount := 0
// 	curStr := ""

// 	for  i:=2 ; i <= n; i++{
// 		// 一次遍历n
// 		for _, v := range prevStr{
// 			vChar := string(v)
// 			if curChar == vChar{
// 				// 如果当前字符和正在遍历的字符相等， 则将数量+1
// 				curCount++
// 			}else{
// 				// 如果不相等的话， 比如 21, curCount 为1， curChar为2
// 				curStr += fmt.Sprintf("%d%s", curCount, curChar)
// 				// 开始遍历新的字符
// 				curChar = vChar
// 				curCount = 1
// 			}
// 		}
// 		prevStr = curStr + fmt.Sprintf("%d%s", curCount, curChar)
// 		curStr = ""
// 		curCount = 0
// 		curChar = string(prevStr[0])
// 		fmt.Printf("%s \n",prevStr)
// 	}

// 	return prevStr
// }

func countAndSay(n int) string {
	// 递归解法
	if n == 1{
		return "1"
	}
	str := countAndSay(n-1)
	curChar := string(str[0])  // 当前比较的字符串
	curCount := 0  // 当前字符的计数
	r := ""  // 返回的结果
	
	fmt.Println(str, n, r)

	for _, v := range str{
		vChar := string(v)
		if curChar == vChar{
			curCount ++
		}else{
			r += fmt.Sprintf("%d%s", curCount, curChar)
			curChar = vChar
			curCount = 1
		}
	}
	r = r + fmt.Sprintf("%d%s", curCount, curChar)
	return r

}

// @lc code=end
