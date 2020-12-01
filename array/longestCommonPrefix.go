package main

import (
	"fmt"
	"strings"
)

/*
	如果strings.Index(x1,x) == 0，则直接跳过（因为此时x就是x1的最长公共前缀），对比下一个元素。（如flower和flow进行比较）
	如果strings.Index(x1,x) != 0, 则截取掉基准元素x的最后一个元素，再次和x1进行比较，直至满足string.Index(x1,x) == 0，此时截取
	后的x为x和x1的最长公共前缀。（如flight和flow进行比较，依次截取出flow-flo-fl，直到fl被截取出，此时fl为flight和flow的最长公共前缀）
*/

func longestCommonPrefix(strs []string) string {
	//if len(strs) < 0 || strs[0] == "" {
	//	return ""
	//}
	//prefix := strs[0]
	//for _, v := range strs {
	//	for strings.Index(v, prefix) != 0 {
	//		prefix = prefix[:len(prefix)-1]
	//	}
	//}
	//return prefix

	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs[1:] {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() { 
	stris := []string{"oi", "flow", "flight"}
	fmt.Println(longestCommonPrefix(stris))
}
