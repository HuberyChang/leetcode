package main

import (
	"fmt"
)

func quicksort(t []int, low, high, k int) int {
	if low == high {
		return t[low]
	}
	temp := t[low]
	p, q := low, high
	for {
		if p < q {
			for {
				if p < q && t[q] >= temp {
					q--
				} else {
					t[p] = t[q]
					break
				}
			}
			for {
				if p < q && t[p] < temp {
					p++
				} else {
					t[q] = t[p]
					break
				}
			}
		} else {
			break
		}
		t[q] = temp
		if q-low+1 >= k {
			quicksort(t, 1, q, k)
		} else {
			quicksort(t, q+1, high, k-(q-low+1))
		}
	}
	return t[k]
}

func main() {
	//var a, b int
	//fmt.Scanf("%d %d\n", &a, &b)
	//q := make([]int, a)
	//for i := 0; i < a; i++ {
	//	fmt.Scanf("%d", &q[i])
	//}
	q := []int{2, 1, 52, 13, 4}
	fmt.Println(quicksort(q, 0, len(q)-1, 2))
}
