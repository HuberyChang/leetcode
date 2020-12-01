package main

import "fmt"

/*
	基数排序：
		对大范围的数据，例如手机号
		字符串
*/
const AsciiRange = 128

func radixSort(arr []string, maxlength int) []string { // maxlength是字符串数组中最长字符串的长度

	//排序结果数组，用于存储每一次按位排序的临时结果
	sortedArray := make([]string, len(arr))

	//从字符串数组中每个字符串的个位开始比较，一直比较到最长字符串的最高位。
	for t := maxlength - 1; t >= 0; t-- {
		//计数排序的过程，分成三步：

		//1.创建辅助排序的统计数组，并把待排序的字符对号入座，
		//这里为了代码简洁，直接使用ascii码范围作为数组长度
		count := make([]int, AsciiRange)
		for i := 0; i < len(arr); i++ { // 对字符串数组中每个字符串的对应下标的字符进行统计，个位开始
			index := getCharIndex(arr[i], t)
			count[index]++
		}

		//2.统计数组做变形，后面的元素等于前面的元素之和 
		for i := 1; i < len(count); i++ {
			count[i] = count[i] + count[i-1]
		}

		//3.倒序遍历原始数列，从统计数组找到正确位置，输出到结果数组
		for i := len(arr) - 1; i >= 0; i-- {
			index := getCharIndex(arr[i], t)
			sortedIndex := count[index] - 1
			sortedArray[sortedIndex] = arr[i]
			count[index]--
		}
		//下一轮排序需要以上一轮的排序结果为基础，因此把结果复制给array
		copy(arr, sortedArray)
	}
	return arr
}

//获取字符串第k位字符所对应的ascii码序号
func getCharIndex(str string, k int) int {
	//如果字符串长度小于k，直接返回0，相当于给不存在的位置补0
	if len(str) < k+1 {
		return 0
	}
	return int(str[k])
}

func main() {
	list := []string{"banana", "apple", "orange", "ape", "he", "bb"}
	fmt.Println(radixSort(list, 6))
}
