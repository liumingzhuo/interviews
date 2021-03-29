//判断两个给定的字符串排序后是否⼀致
//给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀
//个字符串。 这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。给定⼀个string s1和⼀个string s2，请返回⼀个bool，代表两串是否重新排列后可相同。 保证两串的
//⻓度都⼩于等于5000。
package diffregroupstr

import "strings"

func isGrouStr(str1, str2 string) bool {
	//使用strings.Count() 如果每个字符个数都相同 ，那么排序后一定一致
	s1 := []rune(str1)
	s2 := []rune(str2)
	if len(s1) > 5000 || len(s2) > 5000 || len(s1) != len(s2) {
		return false
	}
	for _, s := range s1 {
		if strings.Count(str1, string(s)) != strings.Count(str2, string(s)) {
			return false
		}
	}
	return true
}
