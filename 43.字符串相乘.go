/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (42.61%)
 * Likes:    418
 * Dislikes: 0
 * Total Accepted:    81.8K
 * Total Submissions: 190K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 *
 * 示例 1:
 *
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 *
 * 示例 2:
 *
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 *
 * 说明：
 *
 *
 * num1 和 num2 的长度小于110。
 * num1 和 num2 只包含数字 0-9。
 * num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 *
 *
 */

// @lc code=start

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	// 两个字符串相乘， 最多可能得到 m+n位数
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 两个元素相乘的结果， 与res[i+j] 和 res[i+j+1] 有关
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	fmt.Println(res)

	i := 0
	for ; i < len(res); i++ {
		if res[i] != 0 {
			break
		}
	}

	str := ""
	for ; i < len(res); i++ {
		str += strconv.Itoa(res[i])
	}
	if str == ""{
		return "0"
	}

	return str
}

func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2) // 分别代表两个字符串的长度
	ans := "0"                   // 最终返回的结果
	for i := m - 1; i >= 0; i-- {
		cur := "" // 当前相乘后所得到的结果
		// 将0添加到cur的末尾
		for j := m - 1; j > i; j-- {
			cur += "0"
		}
		add := 0 // 进位
		// num1当前元素的值
		y := int(num1[i] - '0')

		// 编辑num2中每个元素的值进行相乘
		for j := n - 1; j >= 0; j-- {
			// num2当前元素的值
			x := int(num2[j] - '0')
			//两数相乘后的结果
			product := x*y + add
			// 将相乘后的结果添加到cur中
			cur = strconv.Itoa(product%10) + cur
			add = product / 10
		}
		// 将add多出来的数添加到cur的前面
		for add != 0 {
			cur = strconv.Itoa(add%10) + cur
			add /= 10
		}
		ans = addStrings(ans, cur)
	}
	return ans
}

// 两个字符串相加
func addStrings(num1, num2 string) string {
	add := 0  // 负责进位
	ans := "" //负责返回最终的结果
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		r := x + y + add
		ans = strconv.Itoa(r%10) + ans
		add = r / 10
	}
	return ans
}

// @lc code=end
