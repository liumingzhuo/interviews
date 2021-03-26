//判断字符串中字符是否全都不同
//请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允
//许使⽤额外的存储结构】。 给定⼀个string，请返回⼀个bool值,true代表所有字符全都
//不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的⻓
//度⼩于等于【3000】。

package diffstrings

import "strings"

func DiffStrings(str string) bool {
	if strings.Count(str, "") > 3000 {
		return false
	}
	for _, s := range str {
		if s > 127 {
			return false
		}
		//出现大于1次 就是重复
		if strings.Count(str, string(s)) > 1 {
			return false
		}
	}

	return true
}
