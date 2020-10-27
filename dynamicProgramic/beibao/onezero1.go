package main

import "fmt"

/*
	输入格式
		第一行两个整数，N，V，用空格隔开，分别表示物品数量和背包容积。

		接下来有 N 行，每行两个整数 vi,wi，用空格隔开，分别表示第 i 件物品的体积和价值。

	输出格式
		输出一个整数，表示最大价值。

	数据范围
		0<N,V≤1000
		0<vi,wi≤1000
	输入样例
		4 5
		1 2
		2 4
		3 4
		4 5
	输出样例：
		8
*/

const N = 1010

func knapsack() int {
	var (
		n, m int
	)
	fmt.Scanf("%d %d\n", &n, &m)
	v := make([]int, N)
	w := make([]int, N)
	f := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d %d\n", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			f[j] = max1(f[j], f[j-v[i]]+w[i])
		}
	}
	return f[m]
}

func max1(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	ret := knapsack() 
	fmt.Println(ret)
}
