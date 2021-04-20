//反转字符串
//请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
//给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

package reversestr

func reverseString(str string) (string, bool) {
	s := []rune(str)
	l := len(s)
  if l > 5000{
    return str, false
  }
  for i := 0; i < l/2; i++{
    s[i], s[l-i-1] = s[l-i-1],s[i]
  }
  return string(s), true
}
