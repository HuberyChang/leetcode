package main

import "fmt"

/*
	插入排序的改进型，直接插入排序理解为间隔为1的希尔排序
*/

func shellSort(arr []int) []int {

	//折半间隔
	for gap := len(arr) >> 1; gap > 0; gap /= 2 {
		// i从下标大小为gap的位置开始，一直到数组长度为止。
		for i := gap; i < len(arr); i++ {
			fmt.Println(i, "+++++++++++")
			//i,j同时指向摸得牌
			for j := i; j > gap-1; j -= gap { // gap最小为1，j的最小值为j=i=gap，也为1，插入排序中j>0,所以此处j>gap-1。
				// j-=gap 是因为要选相同步长的元素作为子序列，然后执行插排操作
				fmt.Println(i, "==============", j)
				//摸得牌比现在（现在的牌已经是排序好的）最大的牌小，就交换 
				if arr[j] < arr[j-gap] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
				}
			}
		}
	}

	//knuth间隔

	/*	h := 1
		for h <= len(arr)/3 {
			h = h*3 + 1
		}
		for gap := h; gap > 0; gap = (gap - 1) / 3 {
			for i := gap; i < len(arr); i++ {
				fmt.Println(i, "+++++++")
				//i,j同时指向摸得牌
				for j := i; j > gap-1; j -= gap {
					fmt.Println(i, "=======", j)
					//摸得牌比现在（现在的牌已经是排序好的）最大的牌小，就交换
					if arr[j] < arr[j-gap] {
						arr[j], arr[j-gap] = arr[j-gap], arr[j]
					}
				}
			}
		}*/

	return arr
}

func main() {
	list := []int{3, 9, 8, 1, 6}
	fmt.Println(shellSort(list))
}
