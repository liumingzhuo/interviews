//字符串替换问题
//请编写一个方法，将字符串中的空格全部替换为“%20”。 假定该字符串有足够的空间存放新增的字符，
//并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。
//给定一个string为原始的串，返回替换后的string。
package replacestr

import (
	"errors"
	"strings"
	"unicode"
)

func replaceStr(str string) (string, error) {
	if len([]rune(str)) > 1000 {
		return str, errors.New(" str too long ")
	}
	for _, s := range str {
		//即不是空格  也不是字符
		if string(s) != " " && !unicode.IsLetter(s) {
			return str, errors.New("str is not letter")
		}
	}
	return strings.Replace(str, " ", "%20", -1), nil
}
