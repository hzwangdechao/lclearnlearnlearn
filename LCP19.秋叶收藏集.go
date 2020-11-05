package main

/*
	https://leetcode-cn.com/problems/UlBDOe/
*/

const inf = math.MaxInt32 // 或 math.MaxInt64

func minimumOperations(leaves string) int {
    n := len(leaves)
    f := make([][3]int, n)
    f[0][0] = boolToInt(leaves[0] == 'y')
    f[0][1] = inf
    f[0][2] = inf
    f[1][2] = inf
    for i := 1; i < n; i++ {
        isRed := boolToInt(leaves[i] == 'r')
        isYellow := boolToInt(leaves[i] == 'y')
        f[i][0] = f[i-1][0] + isYellow
        f[i][1] = min(f[i-1][0], f[i-1][1]) + isRed
        if i >= 2 {
            f[i][2] = min(f[i-1][1], f[i-1][2]) + isYellow
        }
    }
    return f[n-1][2]
}


// 动态规划的方法
func method1(leaves string) int{
	n := len(leaves)
	dp := make([][3]int, n)
	dp[0][0] = boolToInt(leaves[0] == 'y')
	dp[0][1] = inf
	dp[0][2] = inf
	dp[1][2] = inf

	for i:=1; i<n; i++{
		isRed := boolToInt(leaves[i] == 'r')
		isYellow := boolToInt(leaves[i] == 'y')

		dp[i][0] = dp[i-1][0] + isYellow
		dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + isRed
		if i >= 2{
			dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + isYellow
		}
	}
	return dp[i-1][2]
}

// 如果是相关颜色的话则返回1， 否则返回0
func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
