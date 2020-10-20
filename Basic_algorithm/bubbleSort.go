package main

import "fmt"

/*
	n个数比较n-1轮，每轮的次数就是：n-轮数
*/

func bubleSort(arr []int) []int {
	length := len(arr)
	//i作为比较轮数
	for i := 0; i < length-1; i++ {
		//j作为每轮比较次数
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	list := []int{1, 8, 4, 9, 2, 5, 3}
	fmt.Println(len(list))
	fmt.Println(bubleSort(list))
}
