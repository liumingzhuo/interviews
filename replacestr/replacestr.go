//字符串替换问题
package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func replaceStr(str string) (string, error) {
	if len([]rune(str)) > 1000 {
		return str, errors.New("str is too long")
	}
	//unicode.IsLetter 判断字符是否是字⺟
	for _, v := range str {
		if string(v) != " " && !unicode.IsLetter(v) {
			return str, errors.New("str has non letter")
		}
	}
	return strings.Replace(str, " ", "%20", -1), nil
}

func main() {
	str, err := replaceStr("abca *&q")
	if err != nil {
		fmt.Println(err)
		return
	}
	print(str)

}
