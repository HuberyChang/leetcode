package main

import (
	"fmt"
	"strings"
)

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
