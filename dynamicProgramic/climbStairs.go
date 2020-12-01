package main

import "fmt"

/*
	dp[n] 表示能到达第 n 阶的方法总数 
*/

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	/*
		dp[0] = 0，根据转移方程，dp[2]=1，但是dp[2]=2，初始值不对。所以dp[2]也是一个初始值
	*/
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(10))
}
