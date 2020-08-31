package main

import "fmt"

const P = 110

func fenzu() int {
	var p1, q1, s1 int
	ff := make([]int, P)
	vv := make([]int, P)
	ww := make([]int, P)
	fmt.Scanf("%d %d\n", &p1, &q1)
	for i := 0; i < p1; i++ {
		fmt.Scanf("%d\n", &s1)
		for j := 0; j < s1; j++ {
			fmt.Scanf("%d %d\n", &vv[j], &ww[j])
		}
		for j := q1; j >= 0; j-- {
			for k := 0; k < s1; k++ {
				ff[j] = max4(ff[j], ff[j-vv[k]]+ww[k])
			}
		}
	}
	return ff[q1]
}

func max4(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(fenzu())
}
