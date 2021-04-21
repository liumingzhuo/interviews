/*
机器人坐标问题
有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。 可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。 问最后机器人的坐标是多少？
*/
package robot

import (
	"unicode"
)

//定义上下左右
const (
	Left = iota
	Top
	Right
	Bottom
)

//move  x 关注左右  y关注上下 首次z0 为top
func move(cmd string, x0, y0, z0 int) (x, y, z int) {
	//起点是0.0 方向为Y   那么左转一次是left
	//右转一次是right
	x, y, z = x0, y0, z0
	repeat := 0
	cmdStr := ""
	for _, s := range cmd {
		switch {
		case unicode.IsNumber(s):
			repeat = repeat*10 + (int(s) - '0') //计算重复次数
		case repeat > 0 && s != '(' && s != ')':
			cmdStr = cmdStr + string(s) // 拼接括号内的字符
		case s == ')': //括号结束后 递归调用括号内的步数
			for i := 0; i < repeat; i++ {
				x, y, z = move(cmdStr, x, y, z)
			}
			repeat = 0
			cmdStr = ""
		case s == 'L':
			z = (z + 1) % 4 //逆方向加1
		case s == 'R':
			z = (z - 1 + 4) % 4 //正向+1
		case s == 'F':
			switch {
			case z == Left || z == Right:
				x = x - z + 1 //当z 为0时说明沿x轴正向+1  z为2时延x轴负向-1
			case z == Top || z == Bottom:
				y = y - z + 2 //当z 为1时说明沿y轴正向+1  z为3时沿y轴负向-1
			}
		case s == 'B':
			switch {
			case z == Left || z == Right:
				x = x + z - 1
			case z == Top || z == Bottom:
				y = y + z - 2
			}
		}

	}
	return
}
