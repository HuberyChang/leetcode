package main

import "fmt"

func selectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	list := []int{3, 2, 5, 7, 9} 
	fmt.Println(selectSort(list))
}
