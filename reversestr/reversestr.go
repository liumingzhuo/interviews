//反转字符串
package reversestr

func reverseString(str string) (string, bool) {
	ss := []rune(str)
	l := len(ss)
	if l > 5000 {
		return str, false
	}
	for i := 0; i < l/2; i++ {
		ss[i], ss[l-1-i] = ss[l-1-i], ss[i]
	}
	return string(ss), true
}
