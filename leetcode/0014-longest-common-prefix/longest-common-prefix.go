package leetcode

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string{
	target := strs[0]
	prefix := ""
	var stop bool
	for _,item := range strings.Split(target,"") {
		prefix += item
		for _,str := range strs {
			if !strings.HasPrefix(str, prefix) {
				stop = true
				break
			}
		}
		if stop {
			prefix = prefix[0:len(prefix)-1]
			break
		}
	}
	return prefix
}
