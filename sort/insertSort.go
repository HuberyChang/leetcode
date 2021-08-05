package main

import "fmt"

func insertSort(arr []int) []int {
	//总共len（arr）张牌。从第二张牌开始比较。
	for i := 1; i < len(arr); i++ {
		//将第i张牌插到前面。j是同时指向第i张牌的下标，从第二张牌开始才有插入操作，所以j最小取1，即j>0
		for j := i; j > 0; j-- {
			//第i张牌和前面的有序数组中元素依次比较，判断大小并交换 
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	list := []int{4, 2, 3, 6, 9}
	fmt.Println(insertSort(list))
}
