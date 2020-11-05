*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (54.27%)
 * Likes:    869
 * Dislikes: 0
 * Total Accepted:    161.3K
 * Total Submissions: 292.9K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start

// 按键对应的字母
var phoneMap map[string]string = map[string]string{
	"2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0{
		return res
	}
	var dfs func(index int, str string)
	dfs = func (index int, str string)  {
		if index == len(digits){
			res = append(res, str)
			return
		}

		cur := string(digits[index])
		letter := phoneMap[cur]
		letterCount := len(letter)

		for i:=0; i<letterCount; i++{
			dfs(index+1, str+string(letter[i]))
		}
	}
	dfs(0, "")
	return res
}

// // 存放最终的结果
// var res []string

// func letterCombinations(digits string) []string {
// 	if len(digits) == 0{
// 		return []string{}
// 	}

// 	res = []string{}

// 	traceBack(digits, 0, "")

// 	return res
// }

// func traceBack(digits string, index int, str string)  {
// 	if index == len(digits){
// 		res = append(res, str)
// 		return
// 	}
// 	// 当前数字
// 	curNum := string(digits[index])
// 	// 当前数字对应的字符
// 	numLetters := phoneMap[curNum]
// 	// 按键上字符的长度
// 	letterCount := len(numLetters)

// 	for i:=0; i<letterCount; i++{
// 		traceBack(digits, index+1, str+string(numLetters[i]))
// 	}
// }



// @lc code=end
