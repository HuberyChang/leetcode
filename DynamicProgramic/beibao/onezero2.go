package main

import "fmt"

/*

 */

func knapsack1(w, v []int, total int) int {
	lens := len(w)
	if lens == 0 {
		return 0
	}
	space := make([]int, total+1)
	if w[0] <= total {
		for i := w[0]; i <= total; i++ {
			space[i] = v[0]
		}
	}
	for i := 1; i < lens; i++ {
		for j := total; j >= w[i]; j-- {
			space[j] = maax(space[j], space[j-w[i]]+v[i])
		}
	}
	return space[total]
}

func maax(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func main() {
	w := []int{1, 2, 3}
	v := []int{6, 10, 12}
	fmt.Println(knapsack1(w, v, 5))
}
