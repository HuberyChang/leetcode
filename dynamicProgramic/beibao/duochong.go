package main

import "fmt"

/*
	有s+1种选法，0,1,2，...,s。是01背包的扩展，所以枚举还是得从大到小

	# 0 < k && k * w[i] <= j &&  k <= num[i]
	dp[i][j] = max(dp[i-1][j-k*w[i]] + k*v[i], dp[i-1][j])
	dp[i][j]：表示前i件物品的在体积为j时的最大价值 
*/

const M = 1000

func duochong() int {
	var p, q int
	fmt.Scanf("%d %d\n", &p, &q)
	v2 := make([]int, M)
	w2 := make([]int, M)
	f2 := make([]int, M)
	s2 := make([]int, M)
	for i := 1; i <= p; i++ {
		fmt.Scanf("%d %d %d\n", &v2[i], &w2[i], &s2[i])
		for j := q; j >= 0; j-- {
			for k := 1; k <= s2[i] && k*v2[i] <= j; k++ {
				f2[j] = max3(f2[j], f2[j-k*v2[i]]+k*w2[i])
			}
		}
	}
	return f2[q]
}

func max3(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(duochong())
}
