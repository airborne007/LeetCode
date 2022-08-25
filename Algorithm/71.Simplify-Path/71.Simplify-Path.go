package simplifypath

import (
	"strings"
)

func simplifyPath(path string) string {
	stk := []string{}
	for _, str := range strings.Split(path, "/") {
		if str == "" || str == "." {
			continue
		}
		if str == ".." {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, str)
		}
	}
	return "/" + strings.Join(stk, "/")
}
