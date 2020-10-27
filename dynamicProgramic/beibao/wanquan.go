package main

import "fmt"

const Nw = 1010

func wanquan() int {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	v1 := make([]int, Nw)
	w1 := make([]int, Nw)
	f1 := make([]int, Nw)
	for i := 1; i <= a; i++ {
		fmt.Scanf("%d %d\n", &v1[i], &w1[i])
	}
	for i := 1; i <= a; i++ {
		for j := v1[i]; j <= b; j++ {
			f1[j] = max2(f1[j], f1[j-v1[i]]+w1[i])
		}
	}
	return f1[b]
}

func max2(a, b int) int {
	if a >= b {
		return a
	} else {
		return b 
	}
}

func main() {
	fmt.Println(wanquan())
}
