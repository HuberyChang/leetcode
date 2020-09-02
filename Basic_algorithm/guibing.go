package main

import "fmt"

func merge_sort(l []int, low, high int) []int {
	if low >= high {
		return nil
	}
	mid := (low + high) >> 1
	merge_sort(l, low, mid)
	merge_sort(l, mid+1, high)
	k := 0
	i, j := low, mid+1
	tmp := make([]int, len(l))
	for {
		if i <= mid && j <= high {
			if l[i] <= l[j] {
				tmp[k] = l[i]
				i++
				k++
			} else {
				tmp[k] = l[j]
				k++
				j++
			}
		} else {
			break
		}
	}
	if i <= mid {
		tmp[k] = l[i]
		k++
		i++
	}
	if j <= high {
		tmp[k] = l[j]
		k++
		j++
	}
	for i, j := low, 0; i <= high; i++ {
		l[i] = tmp[j]
	}
	return l
}

func main() {
	//var m, n int
	//fmt.Scanf("%d %d\n", &m, &n)
	ll := []int{2, 3, 1, 5, 4}
	rete := merge_sort(ll, 0, len(ll)-1)
	fmt.Println(rete)
}
