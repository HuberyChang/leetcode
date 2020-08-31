package main

import "fmt"

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
