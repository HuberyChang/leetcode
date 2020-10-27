package main

import "fmt"

/*
	按照冒泡排序的思想：把相邻的元素两两比较，根据大小来交换元素的位置

	冒泡的规则：

		◆ 每一轮获取第一个数和后面的数据进行依次比较的过程，称为一轮冒泡的过程
		◆ 每一轮冒泡.都是先拿第一个数，依次比对相邻的两个数，如果前一个数比后一个数大，则交换他们的位置，这一轮比较完毕，会把最大的数放在最后面。
		◆ 然后反复重复上面的步骤（每一轮都能将前面数据中一个最大数，放到后面），直到一轮冒泡下来没有任何数据需交互位置,此时数据已经为有序状态

	冒泡的次数：

		◆ 假设列表的长度为n，冒泡排序是每次拿出来第一个元素，需要拿多少次呢？
		应该是列表的长度减1，意味着每一个长度为n的列表，需要冒泡 n-1 次

	每次冒泡比较的次数：

		◆ 第m次冒泡，需要进行依次比较的次数为n-m次，每一次冒泡，都能排好一个数据的顺序，那么随着次的增加排好的数据也会越多，需要比较的数据就越少。

	n个数比较n-1轮，每轮的次数就是：n-轮数
*/

func bubleSort(arr []int) []int {
	//第一版
	/*
		length := len(arr)
		//i作为比较轮数，只有一个数就没有比较，因此比较轮数就从0开始。但是最多只有n-1轮（0到n-2）
		for i := 0; i < length-1; i++ {
			//j作为每轮比较次数。但是每次比较都是从第一个数开始和后面的数据进行比较。
			for j := 0; j < length-1-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	*/

	//第二版
	/*
		for i := 0; i < len(arr)-1; i++ {
			isSorted := true
			for j := 0; j < len(arr)-1-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					isSorted = false
				}
			}
			if isSorted {
				break
			}
		}
	*/

	//第三版

	//记录最后一次交换的位置
	lastExchangeIndex := 0
	//无序数列的边界，每次比较只需要比到这里为止
	sortBorder := len(arr) - 1
	for i := 0; i < len(arr)-1; i++ {
		//有序标记，每一轮的初始是true
		isSorted := true
		for j := 0; j < sortBorder; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//有元素交换，所以不是有序，标记变为false
				isSorted = false
				//把无序数列的边界更新为最后一次交换元素的位置
				//fmt.Println(j, "===============")
				lastExchangeIndex = j
			}
		}
		sortBorder = lastExchangeIndex
		//fmt.Println(lastExchangeIndex)
		if isSorted {
			break
		}
	}
	return arr
}

func main() {
	//list := []int{1, 8, 4, 9, 2, 5, 3}
	list := []int{3, 4, 2, 1, 5, 6, 7, 8}
	fmt.Println(len(list))
	fmt.Println(bubleSort(list))
}
