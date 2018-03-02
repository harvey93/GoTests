package stringUtil

import (
	"strings"
)

func Reverse (str string) string {
	var res string = ""
	arr := strings.Split(str, "")
	for i := len(arr) - 1; i >= 0 ; i-- {
		res += arr[i]
	}
	return res
} 